package main

import (
	log "github.com/Sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{}, clientset kubernetes.Interface)
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// TestHandler is a sample implementation of Handler
type TestHandler struct{}

// Init handles any handler initialization
func (t *TestHandler) Init() error {
	log.Info("TestHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *TestHandler) ObjectCreated(obj interface{}, clientset kubernetes.Interface) {
	log.Info("TestHandler.ObjectCreated")
	log.Info(obj)
	createSecret(clientset)
}

// ObjectDeleted is called when an object is deleted
func (t *TestHandler) ObjectDeleted(obj interface{}) {
	log.Info("TestHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *TestHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("TestHandler.ObjectUpdated")
}

func createSecret(clientset kubernetes.Interface) {
	secretPayload := getSecreteObj()
	clientset.CoreV1().Secrets("default").Create(secretPayload)
	log.Info("secret is successfully created in default namespace")
}

func getSecreteObj() *corev1.Secret {
	var secret *corev1.Secret
	secret = &corev1.Secret{
		Type: corev1.SecretTypeOpaque,
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-secret-harsha-nw",
			Labels: map[string]string{
				"heritage": "component-testing",
				"app":      "super8",
			},
		},
	}
	return secret
}
