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
	service := `{"kind":"Service","apiVersion":"v1","metadata":{"name":"nginx-service-test","labels":{"run":"nginx-service"}},"spec":{"ports":[{"port":80,"protocol":"TCP"}],"selector":{"run":"nginx-service"}}}`
	return &KubeService {
		kubeC: cli,
		yaml : service,
	}
}

func (kubeService *KubeService) Create() {
	handler := "CreateService"
	body := strings.NewReader(kubeService.yaml)
	response, err := kubeService.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	VerifyStatusCode(kubeService.kubeC, handler, response.StatusCode)
}

func (kubeService *KubeService) Get() {
	handler := "GetServices"
	response, err := kubeService.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	VerifyStatusCode(kubeService.kubeC, handler, response.StatusCode)
	handler = "GetService"
	response, err = kubeService.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default","name":"nginx-service-test"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	VerifyStatusCode(kubeService.kubeC, handler, response.StatusCode)
}

func (kubeService *KubeService) Delete() {
	handler := "DeleteService"
	response, err := kubeService.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default","name":"nginx-service-test"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	VerifyStatusCode(kubeService.kubeC, handler, response.StatusCode)
}