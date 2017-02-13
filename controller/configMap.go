package controller

// import (
// 	"strings"
// 	"fmt"
// 	"github.com/yangliucheng/easy_http"
// )

// type KubeConfigMap struct {
// 	kubeC *KubeClient
// 	yaml string
// }

// func NewKubeService(cli *KubeClient) *KubeConfigMap {
// 	configMap := `{"kind":"Service","apiVersion":"v1","metadata":{"name":"nginx-service-test","labels":{"run":"nginx-service"}},"spec":{"ports":[{"port":80,"protocol":"TCP"}],"selector":{"run":"nginx-service"}}}`
// 	return &KubeConfigMap {
// 		kubeC: cli,
// 		yaml : configMap,
// 	}
// }

// func (kubeConfigMap *KubeConfigMap) Create() {
// 	handler := "CreateConfigMap"
// 	body := strings.NewReader(kubeConfigMap.yaml)
// 	response, err := kubeConfigMap.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
// 	if err != nil {
// 		fmt.Println("send request fail:",err)
// 		return
// 	}
// 	VerifyStatusCode(kubeConfigMap.kubeC, handler, response.StatusCode, err)
// }

// func (kubeConfigMap *KubeConfigMap) Get() {
// 	handler := "GetConfigMap"
// 	response, err := kubeConfigMap.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
// 	if err != nil {
// 		fmt.Println("send request fail:",err)
// 		return
// 	}
// 	VerifyStatusCode(kubeConfigMap.kubeC, handler, response.StatusCode, err)
// 	handler = "GetService"
// 	response, err = kubeConfigMap.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default","name":"nginx-service-test"}, nil, nil, "")
// 	if err != nil {
// 		fmt.Println("send request fail:",err)
// 		return
// 	}
// 	VerifyStatusCode(kubeConfigMap.kubeC, handler, response.StatusCode, err)
// }

// func (kubeConfigMap *KubeConfigMap) Delete() {
// 	handler := "DeleteConfigMap"
// 	response, err := kubeConfigMap.kubeC.RequestGen.DoHttpRequest(handler, easy_http.Mapstring{"namespace": "default","name":"nginx-service-test"}, nil, nil, "")
// 	if err != nil {
// 		fmt.Println("send request fail:",err)
// 		return
// 	}
// 	VerifyStatusCode(kubeConfigMap.kubeC, handler, response.StatusCode, err)
// }