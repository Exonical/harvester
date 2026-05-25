package setting

import (
	"encoding/json"
	"fmt"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/settings"
	"github.com/harvester/harvester/pkg/util"
)

const (
	rancherDeploymentName = "rancher"
	tlsIngressSecretName  = "tls-ingress"
)

func (h *Handler) syncSSLCertificate(setting *harvesterv1.Setting) error {
	sslCertificate := &settings.SSLCertificate{}
	value := setting.Value
	if value == "" {
		value = setting.Default
	}
	if err := json.Unmarshal([]byte(value), sslCertificate); err != nil {
		return err
	}
	if sslCertificate.CA == "" && sslCertificate.PublicCertificate == "" && sslCertificate.PrivateKey == "" {
		return h.resetCertificates()
	}

	return h.updateCertificates(sslCertificate)
}

func (h *Handler) resetCertificates() error {
	if err := h.updateIngressDefaultCertificate(util.CattleSystemNamespaceName, util.InternalTLSSecretName); err != nil {
		return err
	}
	return h.redeploySSLCertificateWorkload()
}

func (h *Handler) updateCertificates(sslCertificate *settings.SSLCertificate) error {
	if err := h.updateTLSSecret(sslCertificate.PublicCertificate, sslCertificate.PrivateKey); err != nil {
		return err
	}

	if err := h.updateIngressDefaultCertificate(util.CattleSystemNamespaceName, tlsIngressSecretName); err != nil {
		return err
	}
	return h.redeploySSLCertificateWorkload()
}

func (h *Handler) updateTLSSecret(publicCertificate, privateKey string) error {
	tlsSecret, err := h.secretCache.Get(util.CattleSystemNamespaceName, tlsIngressSecretName)
	if err != nil {
		return err
	}
	toUpdateSecret := tlsSecret.DeepCopy()
	toUpdateSecret.Data["tls.crt"] = []byte(publicCertificate)
	toUpdateSecret.Data["tls.key"] = []byte(privateKey)
	_, err = h.secrets.Update(toUpdateSecret)
	return err
}

// updateGatewayTLSCertificate updates the TLS certificate secret referenced by the Gateway resource
func (h *Handler) updateIngressDefaultCertificate(namespace, secretName string) error {
	secret, err := h.secretCache.Get(namespace, secretName)
	if err != nil {
		return fmt.Errorf("failed to get TLS secret %s/%s for Gateway: %w", namespace, secretName, err)
	}

	if secret.Type != "kubernetes.io/tls" {
		return fmt.Errorf("secret %s/%s is not of type kubernetes.io/tls", namespace, secretName)
	}

	return nil
}

func (h *Handler) redeploySSLCertificateWorkload() error {
	return h.redeployDeployment(util.CattleSystemNamespaceName, rancherDeploymentName)
}
