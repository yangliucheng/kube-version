package controller

import (
	"strings"
	"fmt"
	"github.com/yangliucheng/easy_http"
)

const (
	CREATECONFIGMAP 	string = "CreateNamespacedConfigMap"
	GETCONFIGMAP 		string = "ReadNamespacedConfigMap"
	GETCONFIGMAPS 		string = "ListNamespacedConfigMap"
	UPDATECONFIGMAP 	string = "ReplaceNamespacedConfigMap"
	DELETECONFIGMAP 	string = "DeleteNamespacedConfigMap"
	DELETECONFIGMAPS 	string = "DeletecollectionNamespacedConfigMap"	
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

func (kubeConfigMap *KubeConfigMap) Create(out bool) {
	body := strings.NewReader(kubeConfigMap.yaml)
	response, err := kubeConfigMap.kubeC.RequestGen.DoHttpRequest(CREATECONFIGMAP, easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeConfigMap.kubeC.PrintExcel(response, CREATECONFIGMAP, out)
}

func (kubeConfigMap *KubeConfigMap) Get(out bool) {
	response, err := kubeConfigMap.kubeC.RequestGen.DoHttpRequest(GETCONFIGMAP, easy_http.Mapstring{"namespace": "default","name":"example-config"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeConfigMap.kubeC.PrintExcel(response, GETCONFIGMAP, out)
	response, err = kubeConfigMap.kubeC.RequestGen.DoHttpRequest(GETCONFIGMAPS, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeConfigMap.kubeC.PrintExcel(response, GETCONFIGMAPS, out)
}

func (kubeConfigMap *KubeConfigMap) Put(out bool) {
	body := strings.NewReader(kubeConfigMap.yaml)
	response, err := kubeConfigMap.kubeC.RequestGen.DoHttpRequest(UPDATECONFIGMAP, easy_http.Mapstring{"namespace": "default","name":"example-config"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeConfigMap.kubeC.PrintExcel(response, UPDATECONFIGMAP, out)
}

func (kubeConfigMap *KubeConfigMap) Delete(out bool) {
	response, err := kubeConfigMap.kubeC.RequestGen.DoHttpRequest(DELETECONFIGMAP, easy_http.Mapstring{"namespace": "default","name":"example-config"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeConfigMap.kubeC.PrintExcel(response, DELETECONFIGMAP, out)
}