package controller

import (
	"strings"
	"fmt"
	"github.com/yangliucheng/easy_http"
)

type KubeService struct {
	kubeC *KubeClient
	yaml string
}

func NewKubeService(cli *KubeClient) *KubeService {
	pods := `{"apiVersion":"v1","kind":"Pod","metadata":{"name":"nginx"},"spec":{"containers":[{"name":"nginx","image":"nginx:latest","ports":[{"containerPort":80}]}]}}`
	return &KubeService {
		kubeC: cli,
		yaml : pods,
	}
}

func (kubeService *KubeService) Create() {
	handler := "CreatePods"
	body := strings.NewReader(kubeService.yaml)
	response, err := kubeService.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	VerifyStatusCode(kubeService.kubeC, handler, response.StatusCode)
}

func (kubeService *KubeService) Get() {
	handler := "GetPods"
	response, err := kubeService.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	VerifyStatusCode(kubeService.kubeC, handler, response.StatusCode)
}

func (kubeService *KubeService) Delete() {
	handler := "DeletePods"
	response, err := kubeService.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	VerifyStatusCode(kubeService.kubeC, handler, response.StatusCode)
}