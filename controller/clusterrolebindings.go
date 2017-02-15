package controller

import (
	"strings"
	"fmt"
	"github.com/yangliucheng/easy_http"
)

const (
	 CREATECLUSTERROLEBINDING 	string = "CreateClusterRoleBinding"
	 GETCLUSTERROLEBINDING 	string = "ReadClusterRoleBinding"
	 GETCLUSTERROLEBINDINGS 	string = "ListClusterRoleBinding"
	 UPDATECLUSTERROLEBINDING 	string = "ReplaceClusterRoleBinding"
	 DELETECLUSTERROLEBINDING 	string = "DeleteClusterRoleBinding"
	 DELETECLUSTERROLEBINDINGS string = "DeletecollectionClusterRoleBinding"
)


type KubeClusterRolebinding struct {
	kubeC *KubeClient
	yaml string
}

func NewKubeClusterRolebinding(cli *KubeClient) *KubeClusterRolebinding {
	clusterRole := `{"kind":"RoleBinding","apiVersion":"rbac.authorization.k8s.io/v1alpha1","metadata":{"name":"read-secrets","namespace":"development"},"subjects":[{"kind":"User","name":"dave"}],"roleRef":{"kind":"ClusterRole","name":"secret-reader","apiGroup":"rbac.authorization.k8s.io"}}`
	return &KubeClusterRolebinding {
		kubeC: cli,
		yaml : clusterRole,
	}
}

func (kubeClusterRolebinding *KubeClusterRolebinding) Create(out bool) {

	// create cluster role bind
	kubeClusterRolebinding.kubeC.KubeMap.Segment["clusterrole"].(KubeInter).Create(false)
	body := strings.NewReader(kubeClusterRolebinding.yaml)
	response, err := kubeClusterRolebinding.kubeC.RequestGen.DoHttpRequest(CREATECLUSTERROLEBINDING, easy_http.Mapstring{"namespace": "default"}, body, nil, "")	
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeClusterRolebinding.kubeC.PrintExcel(response, CREATECLUSTERROLEBINDING, out)
	kubeClusterRolebinding.kubeC.KubeMap.Segment["clusterrole"].(KubeInter).Delete(false)
}

func (kubeClusterRolebinding *KubeClusterRolebinding) Get(out bool) {
	response, err := kubeClusterRolebinding.kubeC.RequestGen.DoHttpRequest(GETCLUSTERROLEBINDING, easy_http.Mapstring{"namespace": "default","name":"read-secrets"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeClusterRolebinding.kubeC.PrintExcel(response, GETCLUSTERROLEBINDING, out)
	response, err = kubeClusterRolebinding.kubeC.RequestGen.DoHttpRequest(GETCLUSTERROLEBINDINGS, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeClusterRolebinding.kubeC.PrintExcel(response, GETCLUSTERROLEBINDINGS, out)
}

func (kubeClusterRolebinding *KubeClusterRolebinding) Put(out bool) {
	body := strings.NewReader(kubeClusterRolebinding.yaml)
	response, err := kubeClusterRolebinding.kubeC.RequestGen.DoHttpRequest(UPDATECLUSTERROLEBINDING, easy_http.Mapstring{"namespace": "default","name":"read-secrets"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeClusterRolebinding.kubeC.PrintExcel(response, UPDATECLUSTERROLEBINDING, out)
}

func (kubeClusterRolebinding *KubeClusterRolebinding) Delete(out bool) {
	response, err := kubeClusterRolebinding.kubeC.RequestGen.DoHttpRequest(DELETECLUSTERROLEBINDING, easy_http.Mapstring{"namespace": "default","name":"read-secrets"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeClusterRolebinding.kubeC.PrintExcel(response, DELETECLUSTERROLEBINDING, out)
}