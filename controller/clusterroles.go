package controller

import (
	"strings"
	"fmt"
	"github.com/yangliucheng/easy_http"
)

const (
	 CREATECLUSTERROLE 	string = "CreateClusterRole"
	 GETCLUSTERROLE 	string = "ReadClusterRole"
	 GETCLUSTERROLES 	string = "ListClusterRole"
	 UPDATECLUSTERROLE 	string = "ReplaceClusterRole"
	 DELETECLUSTERROLE 	string = "DeleteClusterRole"
	 DELETECLUSTERROLES string = "DeletecollectionClusterRole"
)


type KubeClusterRole struct {
	kubeC *KubeClient
	yaml string
}

func NewKubeClusterRole(cli *KubeClient) *KubeClusterRole {
	clusterRole := `{"kind":"ClusterRole","apiVersion":"rbac.authorization.k8s.io/v1alpha1","metadata":{"name":"secret-reader"},"rules":[{"apiGroups":[""],"resources":["secrets"],"verbs":["get","watch","list"],"nonResourceURLs":[]}]}`
	return &KubeClusterRole {
		kubeC: cli,
		yaml : clusterRole,
	}
}

func (kubeClusterRole *KubeClusterRole) Create(out bool) {
	body := strings.NewReader(kubeClusterRole.yaml)
	response, err := kubeClusterRole.kubeC.RequestGen.DoHttpRequest(CREATECLUSTERROLE, easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeClusterRole.kubeC.PrintExcel(response, CREATECLUSTERROLE, out)
}

func (kubeClusterRole *KubeClusterRole) Get(out bool) {
	response, err := kubeClusterRole.kubeC.RequestGen.DoHttpRequest(GETCLUSTERROLE, easy_http.Mapstring{"namespace": "default","name":"secret-reader"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeClusterRole.kubeC.PrintExcel(response, GETCLUSTERROLE, out)
	response, err = kubeClusterRole.kubeC.RequestGen.DoHttpRequest(GETCLUSTERROLES, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeClusterRole.kubeC.PrintExcel(response, GETCLUSTERROLES, out)
}

func (kubeClusterRole *KubeClusterRole) Put(out bool) {
	body := strings.NewReader(kubeClusterRole.yaml)
	response, err := kubeClusterRole.kubeC.RequestGen.DoHttpRequest(UPDATECLUSTERROLE, easy_http.Mapstring{"namespace": "default","name":"secret-reader"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeClusterRole.kubeC.PrintExcel(response, UPDATECLUSTERROLE, out)
}

func (kubeClusterRole *KubeClusterRole) Delete(out bool) {
	response, err := kubeClusterRole.kubeC.RequestGen.DoHttpRequest(DELETECLUSTERROLE, easy_http.Mapstring{"namespace": "default","name":"secret-reader"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeClusterRole.kubeC.PrintExcel(response, DELETECLUSTERROLE, out)
}