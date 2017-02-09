package router


import (
	"github.com/yangliucheng/easy_http"
)


var KubeRouter = easy_http.RouterArray{
		{Handler : "CreatePods", Method : "POST", Path : "/api/v1/namespaces/:namespace/pods"},
		{Handler : "GetPods", Method : "GET", Path : "/api/v1/namespaces/:namespace/pods"},
		{Handler : "DeletePods", Method : "DELETE", Path : "/api/v1/namespaces/:namespace/pods"},
}