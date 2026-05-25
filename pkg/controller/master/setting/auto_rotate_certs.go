package setting

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/settings"
	"github.com/harvester/harvester/pkg/util"
)

const (
	defaultReconcilAutoRotateCertsSettingDuration = time.Hour * 24 // 1 day
	certRotationJobName                           = "harvester-cert-rotation"
)

func (h *Handler) syncAutoRotateCerts(setting *harvesterv1.Setting) error {
	logrus.WithFields(logrus.Fields{
		"name":  setting.Name,
		"value": setting.Value,
	}).Info("Processing setting")

	autoRotateCerts := &settings.AutoRotateCerts{}
	if err := json.Unmarshal([]byte(setting.Value), autoRotateCerts); err != nil {
		logrus.WithFields(logrus.Fields{
			"name":  setting.Name,
			"value": setting.Value,
		}).WithError(err).Error("failed to unmarshal setting value")
		return err
	}
	if !autoRotateCerts.Enable {
		return nil
	}

	kubernetesIPs, err := util.GetKubernetesIps(h.endpointCache)
	if err != nil {
		return err
	}
	if len(kubernetesIPs) == 0 {
		err = fmt.Errorf("cluster ip is empty")
		logrus.WithFields(logrus.Fields{
			"name":              setting.Name,
			"service.namespace": metav1.NamespaceDefault,
			"service.name":      "kubernetes",
		}).WithError(err).Error("cluster ip is empty in the endpoints")
		return err
	}

	earliestExpiringCert, err := util.GetAddrsEarliestExpiringCert(kubernetesIPs)
	if err != nil {
		return err
	}
	if earliestExpiringCert == nil {
		logrus.WithFields(logrus.Fields{
			"name":              setting.Name,
			"service.namespace": metav1.NamespaceDefault,
			"service.name":      "kubernetes",
			"reconcileAfter":    defaultReconcilAutoRotateCertsSettingDuration,
		}).Warn("can't find certificate for cluster ip, reconcile setting again")
		h.settingController.EnqueueAfter(setting.Name, defaultReconcilAutoRotateCertsSettingDuration)
		return nil
	}
	logrus.WithField(
		"earliestExpiringCert", earliestExpiringCert,
	).Debug("earliest expiring cert for default/kubernetes ClusterIP")

	expiringInHours := time.Duration(autoRotateCerts.ExpiringInHours) * time.Hour
	if time.Now().Add(expiringInHours).After(earliestExpiringCert.NotAfter) {
		reconcileDuration, err := h.rotateCerts(setting)
		if err != nil {
			return err
		}

		h.settingController.EnqueueAfter(setting.Name, reconcileDuration)
		return nil
	}

	reconcileAfter := defaultReconcilAutoRotateCertsSettingDuration
	if earliestExpiringCert.NotAfter.Sub(time.Now().Add(expiringInHours)) < reconcileAfter {
		reconcileAfter = earliestExpiringCert.NotAfter.Sub(time.Now().Add(expiringInHours))
	}
	logrus.WithFields(logrus.Fields{
		"name":                          setting.Name,
		"expiringInHours":               expiringInHours,
		"earliestExpiringCert.notAfter": earliestExpiringCert.NotAfter,
		"reconcileAfter":                reconcileAfter,
	}).Info("certificate is not expiring, reconcile setting again")
	h.settingController.EnqueueAfter(setting.Name, reconcileAfter)
	return nil
}

func (h *Handler) rotateCerts(setting *harvesterv1.Setting) (time.Duration, error) {
	quickReconcilDuration := time.Second * 30
	jobName := fmt.Sprintf("%s-%d", certRotationJobName, time.Now().Unix())

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobName,
			Namespace: util.HarvesterSystemNamespaceName,
		},
		Spec: batchv1.JobSpec{
			BackoffLimit: ptr.To(int32(3)),
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					RestartPolicy: corev1.RestartPolicyOnFailure,
					HostPID:       true,
					NodeSelector: map[string]string{
						"node-role.kubernetes.io/control-plane": "",
					},
					Tolerations: []corev1.Toleration{
						{
							Key:      "node-role.kubernetes.io/control-plane",
							Operator: corev1.TolerationOpExists,
							Effect:   corev1.TaintEffectNoSchedule,
						},
					},
					Containers: []corev1.Container{
						{
							Name:  "cert-renew",
							Image: "registry.k8s.io/kube-apiserver:v1.31.0",
							Command: []string{
								"/bin/sh", "-c",
								"nsenter -t 1 -m -u -i -n -- kubeadm certs renew all",
							},
							SecurityContext: &corev1.SecurityContext{
								Privileged: ptr.To(true),
							},
						},
					},
				},
			},
		},
	}

	if _, err := h.jobClient.Create(job); err != nil {
		if apierrors.IsAlreadyExists(err) {
			logrus.WithFields(logrus.Fields{
				"name": setting.Name,
				"job":  jobName,
			}).Info("cert rotation job already exists, reconcile later")
			return quickReconcilDuration, nil
		}
		return time.Duration(0), fmt.Errorf("failed to create cert rotation job: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"name": setting.Name,
		"job":  jobName,
	}).Info("cert rotation job created")

	return defaultReconcilAutoRotateCertsSettingDuration, nil
}
