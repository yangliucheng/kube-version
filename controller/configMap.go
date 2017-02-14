package controller

import (
	"strings"
	"fmt"
	"github.com/yangliucheng/easy_http"
)

const (
	create 	string = "CreateNamespacedConfigMap"
	get 	string = "ReadNamespacedConfigMap"
	gets 	string = "ListNamespacedConfigMap"
	put 	string = "ReplaceNamespacedConfigMap"
	delete 	string = "DeletecollectionNamespacedConfigMap"

)

type KubeConfigMap struct {
	kubeC *KubeClient
	yaml string
}

func NewKubeConfigMap(cli *KubeClient) *KubeConfigMap {
	configMap := `{"kind":"ConfigMap","apiVersion":"v1","metadata":{"creationTimestamp":"2016-02-18T19:14:38Z","name":"example-config","namespace":"default"},"data":{"example.property.1":"hello","example.property.2":"world"}}`
	return &KubeConfigMap {
		kubeC: cli,
		yaml : configMap,
	}
}

func (kubeConfigMap *KubeConfigMap) Create() {
	body := strings.NewReader(kubeConfigMap.yaml)
	response, err := kubeConfigMap.kubeC.RequestGen.DoHttpRequest(create, easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeConfigMap.kubeC.PrintExcel(response, create)
}

func (kubeConfigMap *KubeConfigMap) Get() {
	response, err := kubeConfigMap.kubeC.RequestGen.DoHttpRequest(get, easy_http.Mapstring{"namespace": "default","name":"example-config"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeConfigMap.kubeC.PrintExcel(response, get)
	response, err = kubeConfigMap.kubeC.RequestGen.DoHttpRequest(gets, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeConfigMap.kubeC.PrintExcel(response, gets)
}

func (kubeConfigMap *KubeConfigMap) Put() {
	body := strings.NewReader(kubeConfigMap.yaml)
	response, err := kubeConfigMap.kubeC.RequestGen.DoHttpRequest(put, easy_http.Mapstring{"namespace": "default","name":"example-config"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeConfigMap.kubeC.PrintExcel(response, put)
}

func (kubeConfigMap *KubeConfigMap) Delete() {
	response, err := kubeConfigMap.kubeC.RequestGen.DoHttpRequest(delete, easy_http.Mapstring{"namespace": "default","name":"example-config"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeConfigMap.kubeC.PrintExcel(response, delete)
}