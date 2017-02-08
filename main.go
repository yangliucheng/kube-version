package main 

import (
	"flag"
	"kube-version/model"
	"github.com/astaxie/beego"
)
var (
	kube_api = flag.String("kube-api", "conf/kube-api.conf", "kubernetes's api of 1.3.5") 
)

func main() {
	model.Config(*kube_api)
	beego.Run()
}