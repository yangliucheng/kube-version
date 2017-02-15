package controller

import (
	"sync"
	"strings"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"kube-version/model"
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
	RwMutex     *sync.RWMutex
	KubeMap 	*model.OrderMap
	KubeConf 	*KubeConf
	KubeExcel 	*model.KubeExcel
	RequestGen 	*easy_http.RequestGen
}

func newKubeCLient(routerArray easy_http.RouterArray) *KubeClient {

	kubeConf := Config(*kube_conf)
	excel := model.Excel {
		Name : "Name",
		Version : "Version",
		Status 	: "Status",
		Path : "Path",
		Method 	: "Method",
		StatusContent : "StatusContent",
		Case : "Case",
	}
	kubeExcel := model.NewKubeExcel(kubeConf.Results.File, kubeConf.Results.Sheet)
	kubeExcel.Write(kubeConf.Results.File, kubeConf.Results.Sheet, &excel)
	requestGen := easy_http.NewRequestGen(kubeConf.KubeAddr, routerArray)
	rwMutex := new(sync.RWMutex)
	kubeMap := model.NewOrderMap()
	return &KubeClient{
		RwMutex   : rwMutex,
		KubeMap   : kubeMap,
		KubeConf  : kubeConf,
		KubeExcel : kubeExcel,
		RequestGen: requestGen,
	}
}

func Run() {
	kubeClient := newKubeCLient(router.KubeRouter)
	// kubePod := NewKubePod(kubeClient)
	// kubeNamespace := NewKubeNamespace(kubeClient)
	// kubeService := NewKubeService(kubeClient)
	kubeConfigMap := NewKubeConfigMap(kubeClient)
	kubeClient.KubeMap.Set("configmap",kubeConfigMap)
	kubeReplicaset := NewKubeKubeReplicaset(kubeClient)
	kubeClient.KubeMap.Set("replicaset",kubeReplicaset)
	kubeSecret := NewKubeSecret(kubeClient)
	kubeClient.KubeMap.Set("secret",kubeSecret)
	KubeRolebinding := NewKubeRolebinding(kubeClient)
	kubeClient.KubeMap.Set("roleBinding",KubeRolebinding)
	KubeRole := NewKubeRole(kubeClient)
	kubeClient.KubeMap.Set("role",KubeRole)
	KubeClusterRole := NewKubeClusterRole(kubeClient)
	kubeClient.KubeMap.Set("clusterrole",KubeClusterRole)
	for _ , value := range kubeClient.KubeMap.Keys {
		go func(kube KubeInter) {
			kube.Create(true)
			kube.Get(true)
			kube.Put(true)
			kube.Delete(true)
		}(kubeClient.KubeMap.Segment[value].(KubeInter))
	}
}

func (kubeClient *KubeClient) PrintExcel(response *http.Response, handler string, out bool) {

	if !out {
		return
	}

	kubeClient.RwMutex.Lock()

	kubeBody := new(KubeBody)
	byt , _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(byt, kubeBody)
	excel := model.Excel{}
	excel.Name = handler
	excel.Version = kubeClient.KubeConf.KubeVersion
	excel.Status = kubeBody.Status
	if strings.EqualFold(kubeBody.Status, "") {
		excel.Status = "Success"
	}
	excel.Path = response.Request.URL.String()
	excel.Method = response.Request.Method
	excel.StatusContent = response.Status
	excel.Case = kubeBody.Message

	kubeClient.KubeExcel.Write(kubeClient.KubeConf.Results.File, kubeClient.KubeConf.Results.Sheet, &excel)

	if strings.EqualFold(kubeBody.Status, "Failure") {
		kubeClient.KubeExcel.SetFill(kubeClient.KubeConf.Results.File, kubeClient.KubeConf.Results.Sheet, "solid", "00FF0000", "FF000000" , 2)
	}

	kubeClient.RwMutex.Unlock()
}