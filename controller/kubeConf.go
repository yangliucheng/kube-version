package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"flag"
)

var (
	kubeConf = flag.String("kubeConf", "conf/kubernetes.json", "configuration of kubernetes") 
)

type KubeConf struct {
	KubeAddr	string 	`json:"kubeAddr"`
}

func Config(f string) *KubeConf{
	var conf *KubeConf
	byt , err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("read file fail:", err)
		return
	}
	err := json.Unmarshal(byt, conf)
	if err != nil {
		fmt.Println("Unmarshal conf fail:",err)
		return
	}

	return conf
}