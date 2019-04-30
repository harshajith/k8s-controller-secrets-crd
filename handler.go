package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/cbroglie/mustache"
	v1 "github.com/harshajith/k8s-controller-secrets-crd/pkg/apis/gitsecret/v1"
	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{}, clientset kubernetes.Interface)
	ObjectDeleted(obj interface{}, clientset kubernetes.Interface)
	ObjectUpdated(objOld, objNew interface{}, clientset kubernetes.Interface)
}

// TestHandler is a sample implementation of Handler
type GitSecretHandler struct{}

// Init handles any handler initialization
func (t *GitSecretHandler) Init() error {
	log.Info("TestHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *GitSecretHandler) ObjectCreated(obj interface{}, clientset kubernetes.Interface) {
	log.Info("TestHandler.ObjectCreated")
	gitSecret, ok := obj.(*v1.GitSecret)
	if ok {
		secretPayload := getSecreteObj(gitSecret)
		log.Info("Secret payload to be created", *secretPayload)
		clientset.CoreV1().Secrets(gitSecret.ObjectMeta.Namespace).Create(secretPayload)
		log.Info("secret is successfully created in default namespace")
	} else {
		log.Error("Can not cast the object to GitSecret", obj)
	}
}

// ObjectDeleted is called when an object is deleted
func (t *GitSecretHandler) ObjectDeleted(obj interface{}, clientset kubernetes.Interface) {
	log.Info("TestHandler.ObjectDeleted")
	gitSecret, ok := obj.(*v1.GitSecret)
	if ok {
		clientset.CoreV1().Secrets(gitSecret.ObjectMeta.Namespace).Delete(gitSecret.Name, nil)
		log.Info("Successfully deleted a secret", gitSecret)
	} else {
		log.Error("Can not cast the object to GitSecret", obj)
	}

}

// ObjectUpdated is called when an object is updated
func (t *GitSecretHandler) ObjectUpdated(objOld, objNew interface{}, clientset kubernetes.Interface) {
	log.Info("TestHandler.ObjectUpdated")
	gitSecret, ok := objNew.(*v1.GitSecret)
	if ok {
		secretPayload := getSecreteObj(gitSecret)
		log.Info("Secret payload to be updated", *secretPayload)
		clientset.CoreV1().Secrets(gitSecret.ObjectMeta.Namespace).Update(secretPayload)
		log.Info("secret is successfully updated in default namespace")
	} else {
		log.Error("Can not cast the object to GitSecret", objNew)
	}

}

func getSecreteObj(gitSecret *v1.GitSecret) *corev1.Secret {
	log.Info("gitSecret val", gitSecret.Spec.Data)
	templateStr := marshalToYamlStr(&gitSecret.Spec.Data)
	data := enrichTemplateStr(templateStr, gitSecret)

	enrichedDataMap := map[string]string{}
	unmarshalYamlStr(data, &enrichedDataMap)
	secret := populateSecretPayload(enrichedDataMap, gitSecret)
	return secret
}

func populateSecretPayload(enrichedDataMap map[string]string, gitSecret *v1.GitSecret) *corev1.Secret {
	return &corev1.Secret{
		Type: corev1.SecretTypeOpaque,
		ObjectMeta: metav1.ObjectMeta{
			Name: gitSecret.Name,
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

func enrichTemplateStr(templateStr string, gitSecret *v1.GitSecret) string {
	data, err1 := mustache.Render(templateStr, getDataFromConfigServer(gitSecret))
	if err1 != nil {
		log.Panic(err1)
	}
	log.Info("template is: ", data)
	return data
}

func getDataFromConfigServer(gitSecret *v1.GitSecret) map[interface{}]interface{} {
	profile, profileOk := os.LookupEnv("PROFILE")
	project, _ := gitSecret.Spec.Appname
	if !profileOk {
		profile = "default"
	}
	label := gitSecret.Spec.Label
	if len(label) == 0 {
		label = "develop"
	}
	configUrl, ok := os.LookupEnv("CONFIG_SERVER_URL")
	appName := strings.Join([]string{project, "(_)", gitSecret.Spec.Appname, "-", profile, ".yml"}, "")
	if !ok {
		configUrl = "http://config-server:8888/master/git-creds-default.yml"
		log.Info("No config server URL is specified, using the default one http://config-server:8888/master/git-creds-default.yml")
	}
	url := strings.Join([]string{configUrl, label, appName}, "/")
	log.Info("final config-server-url", url)
	resp, err := http.Get(url)
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
