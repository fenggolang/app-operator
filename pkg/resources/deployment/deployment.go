package deployment

import (
	appv1 "github.com/fenggolang/app-operator/pkg/apis/app/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// 新建一个Deployment
func New(app *appv1.App) *appsv1.Deployment {
	labels := map[string]string{"app.example.com/v1": app.Name}
	selector := &metav1.LabelSelector{
		MatchLabels: labels,
	}
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Namespace,
			Labels:    labels,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(app, schema.GroupVersionKind{
					Group:   appv1.SchemeGroupVersion.Group,
					Version: appv1.SchemeGroupVersion.Version,
					Kind:    "App",
				}),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Selector: selector,
			Replicas: app.Spec.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: newContainers(app), // 容器数组，因为一个pod可能多个容器
				},
			},
		},
		Status: appsv1.DeploymentStatus{},
	}
}

func newContainers(app *appv1.App) []corev1.Container {
	var containerPorts []corev1.ContainerPort
	//containerPorts := []corev1.ContainerPort{}

	for _, servicePort := range app.Spec.Ports {
		cPort := corev1.ContainerPort{}
		cPort.ContainerPort = servicePort.TargetPort.IntVal
		containerPorts = append(containerPorts, cPort)
	}
	return []corev1.Container{
		{
			Name:  app.Name,
			Image: app.Spec.Image,
			//Args:            []string{"--bind_ip_all", "--replSet=rs0", "--keyFile=/etc/mongo/default-key"},
			Resources:       app.Spec.Resources,
			ImagePullPolicy: corev1.PullAlways,
			Ports:           containerPorts,
			Env:             app.Spec.Envs,
		},
	}
}
