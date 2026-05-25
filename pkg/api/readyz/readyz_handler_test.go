package readyz

import (
	"testing"

	"github.com/harvester/go-common/common"
	"github.com/harvester/harvester/pkg/generated/clientset/versioned/fake"
	"github.com/harvester/harvester/pkg/util/fakeclients"
	longhornTypes "github.com/longhorn/longhorn-manager/types"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubevirtv1 "kubevirt.io/api/core/v1"
)

func TestClusterReady(t *testing.T) {
	readyCondition := corev1.PodCondition{
		Type:   corev1.PodReady,
		Status: corev1.ConditionTrue,
	}
	notReadyCondition := corev1.PodCondition{
		Type:   corev1.PodReady,
		Status: corev1.ConditionFalse,
	}

	longhornLabels := labels.Set(longhornTypes.GetManagerLabels())
	virtControllerLabels := labels.Set{kubevirtv1.AppLabel: "virt-controller"}

	tests := []struct {
		name                string
		nodes               []*corev1.Node
		pods                []*corev1.Pod
		expectedReady       bool
		expectedMsgContains string
	}{
		{
			name:                "No nodes found",
			nodes:               nil,
			pods:                nil,
			expectedReady:       false,
			expectedMsgContains: "no nodes found",
		},
		{
			name: "Node not ready",
			nodes: []*corev1.Node{
				buildMockNode("node-1", []corev1.NodeCondition{
					{Type: corev1.NodeReady, Status: corev1.ConditionFalse},
				}),
			},
			pods:                nil,
			expectedReady:       false,
			expectedMsgContains: "not all nodes are ready",
		},
		{
			name: "No longhorn manager pods",
			nodes: []*corev1.Node{
				buildMockNode("node-1", []corev1.NodeCondition{
					{Type: corev1.NodeReady, Status: corev1.ConditionTrue},
				}),
			},
			pods:                []*corev1.Pod{},
			expectedReady:       false,
			expectedMsgContains: "longhorn-manager pods not ready",
		},
		{
			name: "Longhorn manager pod not ready",
			nodes: []*corev1.Node{
				buildMockNode("node-1", []corev1.NodeCondition{
					{Type: corev1.NodeReady, Status: corev1.ConditionTrue},
				}),
			},
			pods: []*corev1.Pod{
				buildMockPod(
					"longhorn-manager-1",
					common.LonghornSystemNamespaceName,
					longhornLabels,
					corev1.PodPending,
					[]corev1.PodCondition{notReadyCondition},
				),
			},
			expectedReady:       false,
			expectedMsgContains: "longhorn-manager pods not ready",
		},
		{
			name: "Virt controller pods not ready",
			nodes: []*corev1.Node{
				buildMockNode("node-1", []corev1.NodeCondition{
					{Type: corev1.NodeReady, Status: corev1.ConditionTrue},
				}),
			},
			pods: []*corev1.Pod{
				buildMockPod(
					"longhorn-manager-1",
					common.LonghornSystemNamespaceName,
					longhornLabels,
					corev1.PodRunning,
					[]corev1.PodCondition{readyCondition},
				),
				buildMockPod(
					"virt-controller-1",
					common.HarvesterSystemNamespaceName,
					virtControllerLabels,
					corev1.PodPending,
					[]corev1.PodCondition{notReadyCondition},
				),
			},
			expectedReady:       false,
			expectedMsgContains: "virt-controller pods not ready",
		},
		{
			name: "All components ready - single pods",
			nodes: []*corev1.Node{
				buildMockNode("node-1", []corev1.NodeCondition{
					{Type: corev1.NodeReady, Status: corev1.ConditionTrue},
				}),
			},
			pods: []*corev1.Pod{
				buildMockPod(
					"longhorn-manager-1",
					common.LonghornSystemNamespaceName,
					longhornLabels,
					corev1.PodRunning,
					[]corev1.PodCondition{readyCondition},
				),
				buildMockPod(
					"virt-controller-1",
					common.HarvesterSystemNamespaceName,
					virtControllerLabels,
					corev1.PodRunning,
					[]corev1.PodCondition{readyCondition},
				),
			},
			expectedReady:       true,
			expectedMsgContains: "",
		},
		{
			name: "All components ready - multiple pods with some not ready",
			nodes: []*corev1.Node{
				buildMockNode("node-1", []corev1.NodeCondition{
					{Type: corev1.NodeReady, Status: corev1.ConditionTrue},
				}),
			},
			pods: []*corev1.Pod{
				buildMockPod(
					"longhorn-manager-1",
					common.LonghornSystemNamespaceName,
					longhornLabels,
					corev1.PodRunning,
					[]corev1.PodCondition{readyCondition},
				),
				buildMockPod(
					"longhorn-manager-2",
					common.LonghornSystemNamespaceName,
					longhornLabels,
					corev1.PodPending,
					[]corev1.PodCondition{notReadyCondition},
				),
				buildMockPod(
					"virt-controller-1",
					common.HarvesterSystemNamespaceName,
					virtControllerLabels,
					corev1.PodRunning,
					[]corev1.PodCondition{readyCondition},
				),
				buildMockPod(
					"virt-controller-2",
					common.HarvesterSystemNamespaceName,
					virtControllerLabels,
					corev1.PodFailed,
					[]corev1.PodCondition{notReadyCondition},
				),
			},
			expectedReady:       true,
			expectedMsgContains: "",
		},
		{
			name: "Node ready condition missing",
			nodes: []*corev1.Node{
				buildMockNode("node-1", []corev1.NodeCondition{}),
			},
			pods:                nil,
			expectedReady:       false,
			expectedMsgContains: "not all nodes are ready",
		},
		{
			name: "Multiple nodes all ready",
			nodes: []*corev1.Node{
				buildMockNode("node-1", []corev1.NodeCondition{
					{Type: corev1.NodeReady, Status: corev1.ConditionTrue},
				}),
				buildMockNode("node-2", []corev1.NodeCondition{
					{Type: corev1.NodeReady, Status: corev1.ConditionTrue},
				}),
			},
			pods: []*corev1.Pod{
				buildMockPod(
					"longhorn-manager-1",
					common.LonghornSystemNamespaceName,
					longhornLabels,
					corev1.PodRunning,
					[]corev1.PodCondition{readyCondition},
				),
				buildMockPod(
					"virt-controller-1",
					common.HarvesterSystemNamespaceName,
					virtControllerLabels,
					corev1.PodRunning,
					[]corev1.PodCondition{readyCondition},
				),
			},
			expectedReady:       true,
			expectedMsgContains: "",
		},
		{
			name: "Pod running but missing PodReady condition",
			nodes: []*corev1.Node{
				buildMockNode("node-1", []corev1.NodeCondition{
					{Type: corev1.NodeReady, Status: corev1.ConditionTrue},
				}),
			},
			pods: []*corev1.Pod{
				buildMockPod(
					"longhorn-manager-1",
					common.LonghornSystemNamespaceName,
					longhornLabels,
					corev1.PodRunning,
					[]corev1.PodCondition{readyCondition},
				),
				buildMockPod(
					"virt-controller-1",
					common.HarvesterSystemNamespaceName,
					virtControllerLabels,
					corev1.PodRunning,
					[]corev1.PodCondition{
						{Type: "Initialized", Status: corev1.ConditionTrue},
						{Type: "ContainersReady", Status: corev1.ConditionTrue},
					},
				),
			},
			expectedReady:       false,
			expectedMsgContains: "virt-controller pods not ready",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientset := fake.NewSimpleClientset()

			for _, node := range tt.nodes {
				err := clientset.Tracker().Add(node)
				assert.Nil(t, err, "Mock node should add into fake controller tracker")
			}
			for _, pod := range tt.pods {
				err := clientset.Tracker().Add(pod)
				assert.Nil(t, err, "Mock pod should add into fake controller tracker")
			}

			podCache := fakeclients.PodCache(clientset.CoreV1().Pods)
			nodeCache := fakeclients.NodeCache(clientset.CoreV1().Nodes)

			handler := &ReadyzHandler{
				podCache:  podCache,
				nodeCache: nodeCache,
			}
			ready, msg := handler.clusterReady()
			assert.Equal(t, tt.expectedReady, ready, "Ready status should match expected")
			if tt.expectedMsgContains != "" {
				assert.Contains(t, msg, tt.expectedMsgContains, "Message should contain expected substring")
			} else {
				assert.Empty(t, msg, "Message should be empty when cluster is ready")
			}
		})
	}
}

func buildMockPod(name, namespace string, labels map[string]string, phase corev1.PodPhase, conditions []corev1.PodCondition) *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    labels,
		},
		Status: corev1.PodStatus{
			Phase:      phase,
			Conditions: conditions,
		},
	}
}

func buildMockNode(name string, conditions []corev1.NodeCondition) *corev1.Node {
	return &corev1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Status: corev1.NodeStatus{
			Conditions: conditions,
		},
	}
}
