package controller

import (
	"strings"
	"fmt"
	"github.com/yangliucheng/easy_http"
)

const (
	CREATESECRET 	string = "CreateNamespacedSecret"
	GETSECRET 		string = "ReadNamespacedSecret"
	GETSECRETS 		string = "ListNamespacedSecret"
	UPDATESECRET 	string = "ReplaceNamespacedSecret"
	DELETESECRET 	string = "DeleteNamespacedSecret"
	DELETESECRETS 	string = "DeletecollectionNamespacedSecret"	
)

type KubeSecret struct {
	kubeC *KubeClient
	yaml string
}

func NewKubeSecret(cli *KubeClient) *KubeSecret {
	secret := `{"kind":"Secret","apiVersion":"v1","metadata":{"name":"mysecret"},"type":"Opaque","data":{"username":"YWRtaW4=","password":"MWYyZDFlMmU2N2Rm"}}`
	return &KubeSecret {
		kubeC: cli,
		yaml : secret,
	}
}

func (kubeSecret *KubeSecret) Create(out bool) {
	body := strings.NewReader(kubeSecret.yaml)
	response, err := kubeSecret.kubeC.RequestGen.DoHttpRequest(CREATESECRET, easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeSecret.kubeC.PrintExcel(response, CREATESECRET, out)
}

func (kubeSecret *KubeSecret) Get(out bool) {
	response, err := kubeSecret.kubeC.RequestGen.DoHttpRequest(GETSECRET, easy_http.Mapstring{"namespace": "default","name":"mysecret"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeSecret.kubeC.PrintExcel(response, GETSECRET, out)
	response, err = kubeSecret.kubeC.RequestGen.DoHttpRequest(GETSECRETS, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeSecret.kubeC.PrintExcel(response, GETSECRETS, out)
}

func (kubeSecret *KubeSecret) Put(out bool) {
	body := strings.NewReader(kubeSecret.yaml)
	response, err := kubeSecret.kubeC.RequestGen.DoHttpRequest(UPDATESECRET, easy_http.Mapstring{"namespace": "default","name":"mysecret"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeSecret.kubeC.PrintExcel(response, UPDATESECRET, out)
}

func (kubeSecret *KubeSecret) Delete(out bool) {
	response, err := kubeSecret.kubeC.RequestGen.DoHttpRequest(DELETESECRET, easy_http.Mapstring{"namespace": "default","name":"mysecret"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeSecret.kubeC.PrintExcel(response, DELETESECRET, out)
}