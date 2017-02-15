package controller

import (
	"strings"
	"fmt"
	"github.com/yangliucheng/easy_http"
)

const (
	 CREATEROLE string = "CreateNamespacedRole"
	 GETROLE 	string = "ReadNamespacedRole"
	 GETROLES 	string = "ListNamespacedRole"
	 UPDATEROLE string = "ReplaceNamespacedRole"
	 DELETEROLE string = "DeleteNamespacedRole"
	 DELETEROLES string = "DeletecollectionNamespacedRole"
)


type KubeRole struct {
	kubeC *KubeClient
	yaml string
}

func NewKubeRole(cli *KubeClient) *KubeRole {
	role := `{"kind":"Role","apiVersion":"rbac.authorization.k8s.io/v1alpha1","metadata":{"namespace":"default","name":"pod-reader"},"rules":[{"apiGroups":[""],"resources":["pods"],"verbs":["get","watch","list"]}]}`
	return &KubeRole {
		kubeC: cli,
		yaml : role,
	}
}

func (kubeRole *KubeRole) Create(out bool) {
	body := strings.NewReader(kubeRole.yaml)
	response, err := kubeRole.kubeC.RequestGen.DoHttpRequest(CREATEROLE, easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeRole.kubeC.PrintExcel(response, CREATEROLE, out)
}

func (kubeRole *KubeRole) Get(out bool) {
	response, err := kubeRole.kubeC.RequestGen.DoHttpRequest(GETROLE, easy_http.Mapstring{"namespace": "default","name":"pod-reader"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeRole.kubeC.PrintExcel(response, GETROLE,out)
	response, err = kubeRole.kubeC.RequestGen.DoHttpRequest(GETROLES, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeRole.kubeC.PrintExcel(response, GETROLES,out)
}

func (kubeRole *KubeRole) Put(out bool) {
	body := strings.NewReader(kubeRole.yaml)
	response, err := kubeRole.kubeC.RequestGen.DoHttpRequest(UPDATEROLE, easy_http.Mapstring{"namespace": "default","name":"pod-reader"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeRole.kubeC.PrintExcel(response, UPDATEROLE,out)
}

func (kubeRole *KubeRole) Delete(out bool) {
	response, err := kubeRole.kubeC.RequestGen.DoHttpRequest(DELETEROLE, easy_http.Mapstring{"namespace": "default","name":"pod-reader"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeRole.kubeC.PrintExcel(response, DELETEROLE, out)
}