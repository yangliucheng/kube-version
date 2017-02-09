package controller

import (
	// "fmt"
	"strings"
	"kube-version/router"
	"github.com/yangliucheng/easy_http"
)



type KubeClient struct {
	RequestGen *easy_http.RequestGen
}

func newKubeCLient(kubeAddr string, routerArray easy_http.RouterArray) *KubeClient {

	requestGen := easy_http.NewRequestGen(kubeAddr, routerArray)

	return &KubeClient{
		RequestGen: requestGen,
	}
}

func Run(kubeAddr string) {
	KubeArray := make([]KubeInter,0)
	kubeClient := newKubeCLient(kubeAddr, router.KubeRouter)
	kubePod := NewKubePod(kubeClient)
	KubeArray = append(KubeArray, kubePod)
	
	for _ , kube := range KubeArray {
		go func(kube KubeInter) {
			kube.Create()
			kube.Get()
			kube.Delete()
		}(kube)
	}
}