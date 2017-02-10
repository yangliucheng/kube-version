package controller

import (
	"io/ioutil"
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

func (kubePods *KubePods) Create() {
	body := strings.NewReader(kubePods.Yaml)
	response, err := kubePods.KubeC.RequestGen.DoHttpRequest("CreatePods", easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	byt ,_ := ioutil.ReadAll(response.Body)
	fmt.Println("body:",string(byt))
	// fmt.Println(response.StatusCode)
}

func (kubePods *KubePods) Get() {
	response, err := kubePods.KubeC.RequestGen.DoHttpRequest("GetPods", easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	byt ,_ := ioutil.ReadAll(response.Body)
	fmt.Println("body:",string(byt))
}

func (kubePods *KubePods) Delete() {
	response, err := kubePods.KubeC.RequestGen.DoHttpRequest("DeletePods", easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	fmt.Println(response.StatusCode)
}