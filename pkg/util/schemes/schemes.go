package schemes

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
)

var (
	All                = runtime.NewScheme()
	localSchemeBuilder = runtime.NewSchemeBuilder()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(All))
}

// Register adds a scheme builder function and immediately applies it to the All scheme.
func Register(addToScheme func(*runtime.Scheme) error) error {
	localSchemeBuilder = append(localSchemeBuilder, addToScheme)
	return addToScheme(All)
}

// AddToScheme applies all registered scheme builders to the given scheme.
func AddToScheme(scheme *runtime.Scheme) error {
	return localSchemeBuilder.AddToScheme(scheme)
}
