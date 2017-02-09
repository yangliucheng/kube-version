package main 

import (
	"kube-version/controller"
	"github.com/astaxie/beego"
)

func main() {
	controller.Run()
	beego.Run()
}
