package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AppSpec defines the desired state of App
type AppSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	Replicas  *int32                      `json:"replicas"`           // 副本数
	Image     string                      `json:"image"`              // 镜像
	Resources corev1.ResourceRequirements `json:"resource,omitempty"` // 容器资源限制
	Envs      []corev1.EnvVar             `json:"envs,omitempty"`     // 环境变量,omitempty表示当我们环境变量没有传值的时候，我们在yaml/json中是没有这个字段的
	Ports     []corev1.ServicePort        `json:"ports,omitempty"`    // 端口,为了做端口映射
}

// AppStatus defines the observed state of App
type AppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	/*
		// Status包括２部分：conditions和状Phase
		Conditions []AppConditions
		Phase string
	*/
	// 我们这里直接引用k8s已有的　不用自己再定义
	appsv1.DeploymentStatus `json:",inline"` //inline表示原始显示不会新增字段，如果,inline前面又加了status字段，那么yaml/json中会Status.status.replicas等这样显示
}

/*
type AppConditions struct {
	Type string
	Message string
	Reason string
	Ready bool
	// The last time this condition was updated.
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty" protobuf:"bytes,6,opt,name=lastUpdateTime"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,7,opt,name=lastTransitionTime"`
}
*/

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// App is the Schema for the apps API
// +k8s:openapi-gen=true
type App struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppSpec   `json:"spec,omitempty"`
	Status AppStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppList contains a list of App
type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []App `json:"items"`
}

func init() {
	SchemeBuilder.Register(&App{}, &AppList{})
}
