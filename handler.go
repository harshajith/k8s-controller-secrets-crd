package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/cbroglie/mustache"
	v1 "github.com/harshajith/k8s-controller-secrets-crd/pkg/apis/scbsecret/v1"
	"gopkg.in/yaml.v2"
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
	createSecret(obj, clientset)
}

// ObjectDeleted is called when an object is deleted
func (t *TestHandler) ObjectDeleted(obj interface{}) {
	log.Info("TestHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *TestHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("TestHandler.ObjectUpdated")
}

func createSecret(scbSecret interface{}, clientset kubernetes.Interface) {
	secretPayload := getSecreteObj(scbSecret)
	log.Info("Secret payload to be created", *secretPayload)
	clientset.CoreV1().Secrets("default").Create(secretPayload)
	log.Info("secret is successfully created in default namespace")
}

func getSecreteObj(scbSecret interface{}) *corev1.Secret {
	var secret *corev1.Secret
	original, ok := scbSecret.(*v1.ScbSecret)
	if ok {
		log.Info("original val", original.Spec.Data)
		templateStr := marshalToYamlStr(&original.Spec.Data)
		data := EnrichTemplateStr(templateStr)

		enrichedDataMap := map[string]string{}
		unmarshalYamlStr(data, &enrichedDataMap)
		secret = populateSecretPayload(enrichedDataMap, original.Name)
	}
	return secret
}

func populateSecretPayload(enrichedDataMap map[string]string, name string) *corev1.Secret {
	return &corev1.Secret{
		Type: corev1.SecretTypeOpaque,
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		StringData: enrichedDataMap,
	}
}

func marshalToYamlStr(data *map[string]string) string {
	d, err := yaml.Marshal(data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
	return string(d)
}

func unmarshalYamlStr(data string, enrichedDataMap *map[string]string) {
	err2 := yaml.Unmarshal([]byte(data), enrichedDataMap)
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Info("final map is: ", enrichedDataMap)
}

func EnrichTemplateStr(templateStr string) string {
	data, err1 := mustache.Render(templateStr, getDataFromConfigServer())
	if err1 != nil {
		log.Panic(err1)
	}
	log.Info("template is: ", data)
	return data
}

func getDataFromConfigServer() map[interface{}]interface{} {
	resp, err := http.Get("http://config-server:8888/master/git-creds-default.yml")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Info("response from config server: ", string(body))

	m := make(map[interface{}]interface{})

	err1 := yaml.Unmarshal(body, &m)
	if err1 != nil {
		log.Fatal(err1)
	}

	log.Info("map value foo", m["foo"])
	log.Info("nested value", m["eureka.client.serviceUrl.defaultZone"])

	return m
}
