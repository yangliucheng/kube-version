package controller

import (
	"strings"
	"fmt"
	"github.com/yangliucheng/easy_http"
)

type KubePods struct {
	KubeC *KubeClient
	Yaml string
}

func NewKubePod(cli *KubeClient) *KubePods {
	pods := `{"apiVersion":"v1","kind":"Pod","metadata":{"name":"nginx"},"spec":{"containers":[{"name":"nginx","image":"nginx:latest","ports":[{"containerPort":80}]}]}}`
	return &KubePods {
		KubeC: cli,
		Yaml : pods,
	}
}

func (kubePods *KubePods) Create(handler string) {
	body := strings.NewReader(kubePods.Yaml)
	response, err := kubePods.KubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	VerifyStatusCode(kubePods.KubeC, handler, response.StatusCode)
}

func (kubePods *KubePods) Get(handler string) {
	response, err := kubePods.KubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	VerifyStatusCode(kubePods.KubeC, handler, response.StatusCode)
}

func (kubePods *KubePods) Delete(handler string) {
	response, err := kubePods.KubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	VerifyStatusCode(kubePods.KubeC, handler, response.StatusCode)
}