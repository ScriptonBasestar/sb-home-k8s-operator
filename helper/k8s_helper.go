

func (r *HelloReconciler) createService(m *mygroupv1.Hello) *corev1.Service {

	// Label은 여러곳에서 사용하는 팟의 정보가 담긴 데이터이므로 메서드로 모듈화시켜 정적자원처럼 사용하기 위함입니다.
	myLabel := getLabelForMyCustomResource(m.Name)

	// service struct는 metadata, spec등을 구현할 수 있도록 정의되어있습니다.
	newService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name,
			Namespace: m.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Type:     corev1.ServiceTypeNodePort, // service의 타입은 NodePort입니다.
			Selector: myLabel,
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					NodePort:   31321, // 외부에서 31321포트로 접근하도록 합니다.
					Port:       8375,
					TargetPort: intstr.IntOrString{IntVal: 8395},
				},
			},
		},
	}
	// cr이 삭제됐을때 svc가 남아있는걸 막기위해 ref에 추가
	ctrl.SetControllerReference(m, newService, r.Scheme)
	return newService
}
func (r *HelloReconciler) createDeployment(m *mygroupv1.Hello) *appsv1.Deployment {
	myLabel := getLabelForMyCustomResource(m.Name)

	mySize := m.Spec.Size

	newDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name,
			Namespace: m.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &mySize,
			Selector: &metav1.LabelSelector{
				MatchLabels: myLabel,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: myLabel,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "repo.iris.tools/test/echoproject:4",
						Name:  "echoservice",
					}},
				},
			},
		},
	}

	ctrl.SetControllerReference(m, newDeployment, r.Scheme)
	return newDeployment
}

// pod의 Label은 pod을 identify하는 데이터이므로 메서드로 모듈화시켜 정적자원처럼 사용하기 위함입니다.
func getLabelForMyCustomResource(name string) map[string]string {
	return map[string]string{"app": "echoservice"}
}

