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
	RwMutex       *sync.RWMutex
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
	return &KubeClient{
		RwMutex   : rwMutex,
		KubeConf  : kubeConf,
		KubeExcel : kubeExcel,
		RequestGen: requestGen,
	}
}

func Run() {
	kubeArray := make([]KubeInter, 0)
	kubeClient := newKubeCLient(router.KubeRouter)
	kubePod := NewKubePod(kubeClient)
	kubeNamespace := NewKubeNamespace(kubeClient)
	kubeService := NewKubeService(kubeClient)
	kubeArray = append(kubeArray, kubePod, kubeNamespace,kubeService)

	for _ , value := range kubeArray {
		go func(kube KubeInter) {
			kube.Create()
			kube.Get()
			kube.Delete()
		}(value)
	}
}

func (kubeClient *KubeClient) PrintExcel(response *http.Response, handler string) {

	kubeClient.RwMutex.Lock()

	kubeBody := new(KubeBody)
	byt , _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(byt, kubeBody)
	excel := model.Excel{}
	excel.Name = handler
	excel.Version = kubeClient.KubeConf.KubeVersion
	excel.Status = kubeBody.Status
	if strings.EqualFold(kubeBody.Status, "") {
		excel.Status = "Succeed"
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