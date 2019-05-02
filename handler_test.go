package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	v1 "github.com/harshajith/k8s-controller-secrets-crd/pkg/apis/gitsecret/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var configServerReponse = `com:
  value:
    key: "test-value"`

func TestGetConfigServerURL(t *testing.T) {
	gitSecret := getTestSecret()
	configServerURL := GetConfigServerURLPath(gitSecret)
	equals(t, "develop/fsp(_)secret-manager-app-default.yml", configServerURL)
}

func TestGetSecreteObj(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		equals(t, req.URL.String(), "/develop/fsp(_)secret-manager-app-default.yml")
		rw.WriteHeader(200)
		rw.Write([]byte(configServerReponse))
	}))

	defer server.Close()
	client := &ConfigServerClient{server.Client(), server.URL}

	testGitSecret := getTestSecret()
	createdSecret := GetSecreteObj(testGitSecret, client)
	equals(t, testGitSecret.Name, createdSecret.Name)

	expDataMap := map[string]string{"key": "test-value"}
	equals(t, expDataMap, createdSecret.StringData)
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func getTestSecret() *v1.GitSecret {
	return &v1.GitSecret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "secret-name",
			Namespace: "cicd",
		},
		Spec: v1.GitSecretSpec{
			Appname:      "secret-manager-app",
			Organization: "fsp",
			Label:        "develop",
			Data:         map[string]string{"key": "{{com.value.key}}"},
		},
	}
}

func TestDoStuffWithTestServer(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/some/path")
		// Send response to be tested
		rw.Write([]byte(`Harsha`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := ConfigServerClient{server.Client(), server.URL}

	body, err := DoStuff(&api)

	ok(t, err)
	equals(t, []byte("OK"), body)

}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}
