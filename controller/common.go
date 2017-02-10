package controller

import (
	"fmt"
)

func VerifyStatusCode(kubeC *KubeClient, handler string, statusCode int) {
	if statusCode == 404 {
		router := kubeC.RequestGen.LookUrl(handler)
		fmt.Println(router.Path)
	}
}