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
	KubeMap:= NewOrderMap()
	kubeClient := newKubeCLient(kubeConf.KubeAddr, router.KubeRouter)
	kubePod := NewKubePod(kubeClient)
	KubeMap.Set("CreatePods",kubePod)

	for _ , key := range KubeMap.Keys {
		value := KubeMap.Segment[key]
		go func(kube KubeInter) {
			kube.Create(key.(string))
			kube.Get(key.(string))
			kube.Delete(key.(string))
		}(value.(KubeInter))
	}
}