package controller

import (
	"strings"
	"fmt"
	"github.com/yangliucheng/easy_http"
)

const (
	 CREATEREPLICASET 	string = "CreateNamespacedReplicaSet"
	 GETREPLICASET 		string = "ReadNamespacedReplicaSet"
	 GETREPLICASETS 	string = "ListNamespacedReplicaSet"
	 UPDATEREPLICASET 	string = "ReplaceNamespacedReplicaSet"
	 DELETEREPLICASET 	string = "DeleteNamespacedReplicaSet"
	 DELETEREPLICASETS 	string = "DeletecollectionNamespacedReplicaSet"
)


type KubeReplicaset struct {
	kubeC *KubeClient
	yaml string
}

func NewKubeKubeReplicaset(cli *KubeClient) *KubeReplicaset {
	reslicaset := `{
  "apiVersion": "extensions/v1beta1",
  "kind": "ReplicaSet",
  "metadata": {
    "name": "frontend"
  },
  "spec": {
    "replicas": 3,
    "selector": {
      "matchLabels": {
        "tier": "frontend"
      },
      "matchExpressions": [
        {
          "key": "tier",
          "operator": "In",
          "values": [
            "frontend"
          ]
        }
      ]
    },
    "template": {
      "metadata": {
        "labels": {
          "app": "guestbook",
          "tier": "frontend"
        }
      },
      "spec": {
        "containers": [
          {
            "name": "php-redis",
            "image": "redis:latest",
            "resources": {
              "requests": {
                "cpu": "100m",
                "memory": "100Mi"
              }
            },
            "env": [
              {
                "name": "GET_HOSTS_FROM",
                "value": "dns"
              }
            ],
            "ports": [
              {
                "containerPort": 80
              }
            ]
          }
        ]
      }
    }
  }
}`
	return &KubeReplicaset {
		kubeC: cli,
		yaml : reslicaset,
	}
}

func (kubeReplicaset *KubeReplicaset) Create(out bool) {
	body := strings.NewReader(kubeReplicaset.yaml)
	response, err := kubeReplicaset.kubeC.RequestGen.DoHttpRequest(CREATEREPLICASET, easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeReplicaset.kubeC.PrintExcel(response, CREATEREPLICASET, out)
}

func (kubeReplicaset *KubeReplicaset) Get(out bool) {
	response, err := kubeReplicaset.kubeC.RequestGen.DoHttpRequest(GETREPLICASET, easy_http.Mapstring{"namespace": "default","name":"frontend"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeReplicaset.kubeC.PrintExcel(response, GETREPLICASET, out)
	response, err = kubeReplicaset.kubeC.RequestGen.DoHttpRequest(GETREPLICASETS, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeReplicaset.kubeC.PrintExcel(response, GETREPLICASETS, out)
}

func (kubeReplicaset *KubeReplicaset) Put(out bool) {
	body := strings.NewReader(kubeReplicaset.yaml)
	response, err := kubeReplicaset.kubeC.RequestGen.DoHttpRequest(UPDATEREPLICASET, easy_http.Mapstring{"namespace": "default","name":"frontend"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeReplicaset.kubeC.PrintExcel(response, UPDATEREPLICASET, out)
}

func (kubeReplicaset *KubeReplicaset) Delete(out bool) {
	response, err := kubeReplicaset.kubeC.RequestGen.DoHttpRequest(DELETEREPLICASET, easy_http.Mapstring{"namespace": "default","name":"frontend"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeReplicaset.kubeC.PrintExcel(response, DELETEREPLICASET, out)
}