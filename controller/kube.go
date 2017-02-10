package controller

import (
	"flag"
	"kube-version/router"
	"github.com/yangliucheng/easy_http"
)

var (
	kube_conf = flag.String("kube_conf", "conf/kubernetes.json", "configuration of kubernetes") 
)

func init() {
	flag.Parse()
}

type KubeClient struct {
	RequestGen *easy_http.RequestGen
}

func newKubeCLient(kubeAddr string, routerArray easy_http.RouterArray) *KubeClient {

	requestGen := easy_http.NewRequestGen(kubeAddr, routerArray)

	return &KubeClient{
		RequestGen: requestGen,
	}
}

func Run() {
	kubeConf := Config(*kube_conf)
	KubeArray := make([]KubeInter,0)
	kubeClient := newKubeCLient(kubeConf.KubeAddr, router.KubeRouter)
	kubePod := NewKubePod(kubeClient)
	KubeArray = append(KubeArray, kubePod)

	for _ , kube := range KubeArray {
		go func(kube KubeInter) {
			// kube.Create()
			// kube.Get()
			kube.Delete()
		}(kube)
	}
}