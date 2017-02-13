package controller

import (
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
	KubeExcel 	*model.KubeExcel
	RequestGen 	*easy_http.RequestGen
}

func newKubeCLient(kubeAddr string, routerArray easy_http.RouterArray) *KubeClient {
	excel := model.Excel {
		Name : "Name",
		Version : "Version",
		Status 	: "Status",
		Path : "Path",
		Method 	: "Method",
		StatusContent : "StatusContent",
		Case : "Case",
	}
	path := "/home/yang/test.xlsx"
	sheet := "kube"
	kubeExcel := model.NewKubeExcel(path, sheet)
	kubeExcel.Write(path,sheet, &excel)
	requestGen := easy_http.NewRequestGen(kubeAddr, routerArray)
	return &KubeClient{
		KubeExcel : kubeExcel,
		RequestGen: requestGen,
	}
}

func Run() {
	kubeConf := Config(*kube_conf)
	kubeArray := make([]KubeInter, 0)
	kubeClient := newKubeCLient(kubeConf.KubeAddr, router.KubeRouter)
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

	if response.StatusCode == 404 {
		kubeBody := new(KubeBody)
		byt , _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(byt, kubeBody)
		excel := model.Excel{}
		excel.Name = handler
		excel.Version = "1.3.5"
		excel.Status = kubeBody.Status
		excel.Path = response.Request.URL.String()
		excel.Method = response.Request.Method
		excel.StatusContent = response.Status
		excel.Case = kubeBody.Message
	}
}