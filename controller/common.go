package controller

import (
	"fmt"
)

type KubeBody struct {
	Kind 		string 	`json:"kind"`
	ApiVersion 	string 	`json:"apiVersion"`
	Status 		string 	`json:"status"`
	Message 	string 	`json:"message"`
	Reason 		string 	`json:"reason"`
	Code   		int 	`json:"code"`
}

func VerifyStatusCode(kubeC *KubeClient, handler string, statusCode int, err error) {
	if statusCode == 404 || err != nil {
		router := kubeC.RequestGen.LookUrl(handler)
		fmt.Println(router.Method,router.Path,err)
	}
}