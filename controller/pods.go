package controller

import (
	"strings"
	"fmt"
	"github.com/yangliucheng/easy_http"
)

type KubePods struct {
	kubeC *KubeClient
	yaml string
}

func NewKubePod(cli *KubeClient) *KubePods {
	pods := `{"apiVersion":"v1","kind":"Pod","metadata":{"name":"nginx"},"spec":{"containers":[{"name":"nginx","image":"nginx:latest","ports":[{"containerPort":80}]}]}}`
	return &KubePods {
		kubeC: cli,
		yaml : pods,
	}
}

func (kubePods *KubePods) Create() {
	handler := "CreatePods"
	body := strings.NewReader(kubePods.yaml)
	response, err := kubePods.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	VerifyStatusCode(kubePods.kubeC, handler, response.StatusCode, err)
}

func (kubePods *KubePods) Get() {
	handler := "GetPods"
	response, err := kubePods.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	VerifyStatusCode(kubePods.kubeC, handler, response.StatusCode, err)
}

func (kubePods *KubePods) Delete() {
	handler := "DeletePods"
	response, err := kubePods.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	VerifyStatusCode(kubePods.kubeC, handler, response.StatusCode, err)
}