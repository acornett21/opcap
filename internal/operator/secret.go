package operator

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"
)

func (o operatorClient) CreateSecret(ctx context.Context, name string, content map[string]string, secretType corev1.SecretType, namespace string) (*corev1.Secret, error) {
	secret := corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		StringData: content,
		Type:       secretType,
	}
	err := o.Client.Create(ctx, &secret, &runtimeClient.CreateOptions{})
	if err != nil {
		logger.Errorf("error while creating secret: %s in namespace: %s", name, namespace, err)
		return nil, err
	}

	logger.Debugf("Secret %s created successfully in namespace %s", name, namespace)
	return &secret, nil
}

func (o operatorClient) DeleteSecret(ctx context.Context, name string, namespace string) error {
	secret := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
	logger.Debugf("Deleting secret %s from namespace %s", name, namespace)
	return o.Client.Delete(ctx, &secret, &runtimeClient.DeleteOptions{})
}
