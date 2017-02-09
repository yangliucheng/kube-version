package controller

import (
	"fmt"
	"kube-version/router"
	"github.com/yangliucheng/easy_http"
)



type KubeClient struct {
	RequestGen *easy_http.RequestGen
}

func init() {
	KubeMap := make(easy_http.Handlers,0)
	KubeMap["CreatePods"] = CreatePods
}

func newKubeCLient(kubeAddr string, routerArray easy_http.RouterArray) *KubeClient {

	requestGen := easy_http.NewRequestGen(kubeAddr, routerArray)

	return &KubeClient{
		RequestGen: requestGen,
	}
}

func Run(kubeAddr string) {
	kubeClient := newKubeCLient(kubeAddr, router.KubeRouter)
	for _ , v := range KubeMap {
		go v()
	}
}

func (kube *KubeClient) CreatePods() {
	fmt.Println("CreatePods")
}