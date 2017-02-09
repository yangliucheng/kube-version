package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type KubeConf struct {
	KubeAddr	string 	`json:"kubeAddr"`
}

func Config(f string) *KubeConf{
	conf := new(KubeConf)
	byt , err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("read file fail:", err)
		return conf
	}
	err = json.Unmarshal(byt, conf)
	if err != nil {
		fmt.Println("Unmarshal KubeConf fail:",err)
		return conf
	}
	return conf
}