package router


import (
	"github.com/yangliucheng/easy_http"
)


var KubeRouter = easy_http.RouterArray{
		{Handler : "CreatePods", Method : "POST", Path : "/api/v1/namespaces/:namespace/pods"},
		{Handler : "GetPods", Method : "GET", Path : "/api/v1/namespaces/:namespace/pods"},
		{Handler : "DeletePods", Method : "DELETE", Path : "/api/v1/namespaces/:namespace/pods"},

		{Handler : "CreateNamespace", Method : "POST", Path : "/api/v1/namespaces"},
		{Handler : "GetNamespaces", Method : "GET", Path : "/api/v1/namespaces"},
		{Handler : "GetNamespace", Method : "GET", Path : "/api/v1/namespaces/:name"},
		{Handler : "DeleteNamespace", Method : "DELETE", Path : "/api/v1/namespaces/:name"},

		{Handler : "CreateService", Method : "POST", Path : "/api/v1/namespaces/:namespace/services"},
		{Handler : "GetService", Method : "GET", Path : "/api/v1/namespaces/:namespace/services/:name"},
		{Handler : "GetServices", Method : "GET", Path : "/api/v1/namespaces/:namespace/services"},
		{Handler : "DeleteService", Method : "DELETE", Path : "/api/v1/namespaces/:namespace/services/:name"},
}