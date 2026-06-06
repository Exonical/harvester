package runtime

import (
	"context"

	helmv1 "github.com/k3s-io/helm-controller/pkg/apis/helm.cattle.io/v1"
	cniv1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	lhv1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
	upgradev1 "github.com/rancher/system-upgrade-controller/pkg/apis/upgrade.cattle.io/v1"
	"k8s.io/client-go/rest"

	"github.com/harvester/harvester/pkg/util/crd"
)

// createCRDs creates CRDs needed in integration tests
func createCRDs(ctx context.Context, restConfig *rest.Config) error {
	factory, err := crd.NewFactoryFromClient(ctx, restConfig)
	if err != nil {
		return err
	}
	return factory.
		BatchCreateCRDsIfNotExisted(
			createHelmChartConfigCRD(),
			createNetworkAttachmentDefinitionCRD(),
			createPlanCRD(),
			createHelmChartCRD(),
			createLonghornNodeCRD(),
			createLonghornVolumeCRD(),
			createLonghornReplicaCRD(),
		).
		BatchWait()
}

func createNetworkAttachmentDefinitionCRD() crd.CRD {
	nad := crd.FromGV(cniv1.SchemeGroupVersion, "NetworkAttachmentDefinition", cniv1.NetworkAttachmentDefinition{})
	nad.PluralName = "network-attachment-definitions"
	nad.SingularName = "network-attachment-definition"
	return nad
}

func createHelmChartConfigCRD() crd.CRD {
	mChart := crd.FromGV(helmv1.SchemeGroupVersion, "HelmChartConfig", helmv1.HelmChartConfig{})
	mChart.PluralName = "helmchartconfigs"
	mChart.SingularName = "helmchartconfig"
	return mChart
}

func createPlanCRD() crd.CRD {
	plan := crd.FromGV(upgradev1.SchemeGroupVersion, "Plan", upgradev1.Plan{})
	plan.PluralName = "plans"
	plan.SingularName = "plan"
	return plan
}

func createHelmChartCRD() crd.CRD {
	return crd.NamespacedType("HelmChart.helm.cattle.io/v1").
		WithSchemaFromStruct(helmv1.HelmChart{}).
		WithColumn("Job", ".status.jobName").
		WithColumn("Chart", ".spec.chart").
		WithColumn("TargetNamespace", ".spec.targetNamespace").
		WithColumn("Version", ".spec.version").
		WithColumn("Repo", ".spec.repo").
		WithColumn("HelmVersion", ".spec.helmVersion").
		WithColumn("Bootstrap", ".spec.bootstrap")
}

func createLonghornNodeCRD() crd.CRD {
	return crd.FromGV(lhv1beta2.SchemeGroupVersion, "Node", lhv1beta2.Node{})
}

func createLonghornVolumeCRD() crd.CRD {
	return crd.FromGV(lhv1beta2.SchemeGroupVersion, "Volume", lhv1beta2.Volume{})
}

func createLonghornReplicaCRD() crd.CRD {
	return crd.FromGV(lhv1beta2.SchemeGroupVersion, "Replica", lhv1beta2.Replica{})
}
