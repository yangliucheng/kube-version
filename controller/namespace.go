package controller

import (
	"strings"
	"fmt"
	"github.com/yangliucheng/easy_http"
)

type KubeNamespace struct {
	kubeC *KubeClient
	yaml  string
}

func NewKubeNamespace(cli *KubeClient) *KubeNamespace {
	namespace := `{"apiVersion":"v1","kind":"Namespace","metadata":{"name":"test"}}`
	return &KubeNamespace {
		kubeC: cli,
		yaml : namespace,
	}
}

func (kubeNamespace *KubeNamespace) Create(out bool) {
	handler := "CreateNamespace"
	body := strings.NewReader(kubeNamespace.yaml)
	response, err := kubeNamespace.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeNamespace.kubeC.PrintExcel(response, handler, out)
}

func (kubeNamespace *KubeNamespace) Get(out bool) {
	handler := "GetNamespaces"
	response, err := kubeNamespace.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeNamespace.kubeC.PrintExcel(response, handler, out)
	handler = "GetNamespace"
	response, err = kubeNamespace.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"name":"test"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeNamespace.kubeC.PrintExcel(response, handler, out)
}

func (kubeNamespace *KubeNamespace) Delete(out bool) {
	handler := "DeleteNamespace"
	response, err := kubeNamespace.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"name":"test"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeNamespace.kubeC.PrintExcel(response, handler, out)
}