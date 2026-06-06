package readyz

import (
	"time"

	"github.com/harvester/go-common/common"
	harvesterServer "github.com/harvester/harvester/pkg/server/http"
	longhornTypes "github.com/longhorn/longhorn-manager/types"
	ctlcorev1 "github.com/harvester/harvester/pkg/util/ctlcache/corev1"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubevirtv1 "kubevirt.io/api/core/v1"
)

type ReadyzHandler struct {
	podCache  ctlcorev1.PodCache
	nodeCache ctlcorev1.NodeCache
}

func NewReadyzHandler(podCache ctlcorev1.PodCache, nodeCache ctlcorev1.NodeCache) *ReadyzHandler {
	return &ReadyzHandler{
		podCache:  podCache,
		nodeCache: nodeCache,
	}
}

func (h *ReadyzHandler) Do(ctx *harvesterServer.Ctx) (harvesterServer.ResponseBody, error) {
	ready, msg := h.clusterReady()
	timestamp := time.Now().UTC().Format(time.RFC3339)

	if !ready {
		logrus.Debugf("Cluster not ready: %s", msg)
		ctx.SetStatusOK()
		return map[string]any{
			"ready":     false,
			"msg":       msg,
			"timestamp": timestamp,
		}, nil
	}

	ctx.SetStatusOK()
	return map[string]any{
		"ready":     true,
		"timestamp": timestamp,
	}, nil
}

func (h *ReadyzHandler) clusterReady() (bool, string) {
	if ready, msg := h.nodesReady(); !ready {
		return false, msg
	}

	if ready, msg := h.longhornReady(); !ready {
		return false, msg
	}

	if ready, msg := h.kubevirtReady(); !ready {
		return false, msg
	}

	return true, ""
}

func (h *ReadyzHandler) nodesReady() (bool, string) {
	nodes, err := h.nodeCache.List(labels.Everything())
	if err != nil {
		logrus.Debugf("failed to list nodes: %s", err.Error())
		return false, "failed to list nodes"
	}

	if len(nodes) == 0 {
		return false, "no nodes found"
	}

	for _, node := range nodes {
		ready := false
		for _, cond := range node.Status.Conditions {
			if cond.Type == corev1.NodeReady && cond.Status == corev1.ConditionTrue {
				ready = true
				break
			}
		}
		if !ready {
			return false, "not all nodes are ready"
		}
	}

	return true, ""
}

func (h *ReadyzHandler) longhornReady() (bool, string) {
	longhornManagerSelector := labels.SelectorFromSet(labels.Set(longhornTypes.GetManagerLabels()))
	longhornPods, err := h.podCache.List(common.LonghornSystemNamespaceName, longhornManagerSelector)
	if err != nil {
		logrus.Debugf("failed to check longhorn-manager pods: %s", err.Error())
		return false, "failed to check longhorn-manager pods"
	}

	if !hasAtLeastOneReadyPod(longhornPods) {
		return false, "longhorn-manager pods not ready"
	}

	return true, ""
}

func (h *ReadyzHandler) kubevirtReady() (bool, string) {
	virtControllerSelector := labels.SelectorFromSet(labels.Set{kubevirtv1.AppLabel: "virt-controller"})
	virtPods, err := h.podCache.List(common.HarvesterSystemNamespaceName, virtControllerSelector)
	if err != nil {
		logrus.Debugf("failed to check virt-controller pods: %s", err.Error())
		return false, "failed to check virt-controller pods"
	}

	if !hasAtLeastOneReadyPod(virtPods) {
		return false, "virt-controller pods not ready"
	}

	return true, ""
}

func hasAtLeastOneReadyPod(pods []*corev1.Pod) bool {
	for _, pod := range pods {
		if pod.Status.Phase == corev1.PodRunning && isPodReadyConditionTrue(pod) {
			return true
		}
	}
	return false
}

func isPodReadyConditionTrue(pod *corev1.Pod) bool {
	for _, cond := range pod.Status.Conditions {
		if cond.Type == corev1.PodReady {
			return cond.Status == corev1.ConditionTrue
		}
	}
	return false
}
