package controller

import (
	"strings"
	"fmt"
	"github.com/yangliucheng/easy_http"
)

const (
	 CREATEROLEBINDING 	string = "CreateNamespacedRole"
	 GETROLEBINDING 	string = "ReadNamespacedRoleBinding"
	 GETROLEBINDINGS 	string = "ListNamespacedRole"
	 UPDATEROLEBINDING 	string = "ReplaceNamespacedRoleBinding"
	 DELETEROLEBINDING 	string = "DeleteNamespacedRoleBinding"
	 DELETEROLEBINDINGS string = "DeletecollectionNamespacedRoleBinding"
)


type KubeRolebinding struct {
	kubeC *KubeClient
	yaml string
}

func NewKubeRolebinding(cli *KubeClient) *KubeRolebinding {
	roleBinding := `{"kind":"RoleBinding","apiVersion":"rbac.authorization.k8s.io/v1alpha1","metadata":{"name":"read-pods","namespace":"default"},"subjects":[{"kind":"User","name":"jane"}],"roleRef":{"kind":"Role","name":"pod-reader","apiGroup":"rbac.authorization.k8s.io"}}`
	// clusterRole := `{"kind":"RoleBinding","apiVersion":"rbac.authorization.k8s.io/v1alpha1","metadata":{"name":"read-secrets","namespace":"development"},"subjects":[{"kind":"User","name":"dave"}],"roleRef":{"kind":"ClusterRole","name":"secret-reader","apiGroup":"rbac.authorization.k8s.io"}}`
	return &KubeRolebinding {
		kubeC: cli,
		yaml : roleBinding,
	}
}

func (kubeRolebinding *KubeRolebinding) Create(out bool) {

	// create role binding
	// create role
	kubeRolebinding.kubeC.KubeMap.Segment["role"].(KubeInter).Create(false)
	body := strings.NewReader(kubeRolebinding.yaml)
	response, err := kubeRolebinding.kubeC.RequestGen.DoHttpRequest(CREATEROLEBINDING, easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeRolebinding.kubeC.PrintExcel(response, CREATEROLEBINDING, out)
	kubeRolebinding.kubeC.KubeMap.Segment["role"].(KubeInter).Delete(false)
}

func (kubeRolebinding *KubeRolebinding) Get(out bool) {
	response, err := kubeRolebinding.kubeC.RequestGen.DoHttpRequest(GETROLEBINDING, easy_http.Mapstring{"namespace": "default","name":"read-pods"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeRolebinding.kubeC.PrintExcel(response, GETROLEBINDING, out)
	response, err = kubeRolebinding.kubeC.RequestGen.DoHttpRequest(GETROLEBINDINGS, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeRolebinding.kubeC.PrintExcel(response, GETROLEBINDINGS, out)
}

func (kubeRolebinding *KubeRolebinding) Put(out bool) {
	body := strings.NewReader(kubeRolebinding.yaml)
	response, err := kubeRolebinding.kubeC.RequestGen.DoHttpRequest(UPDATEROLEBINDING, easy_http.Mapstring{"namespace": "default","name":"read-pods"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeRolebinding.kubeC.PrintExcel(response, UPDATEROLEBINDING, out)
}

func (kubeRolebinding *KubeRolebinding) Delete(out bool) {
	response, err := kubeRolebinding.kubeC.RequestGen.DoHttpRequest(DELETEROLEBINDING, easy_http.Mapstring{"namespace": "default","name":"read-pods"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeRolebinding.kubeC.PrintExcel(response, DELETEROLEBINDING, out)
}