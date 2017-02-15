package router


import (
	"github.com/yangliucheng/easy_http"
)


var KubeRouter = easy_http.RouterArray{
		// pod
		// {Handler : "CreatePods", Method : "POST", Path : "/api/v1/namespaces/:namespace/pods"},
		// {Handler : "GetPods", Method : "GET", Path : "/api/v1/namespaces/:namespace/pod"},
		// {Handler : "DeletePods", Method : "DELETE", Path : "/api/v1/namespaces/:namespace/pods"},

		// // namespace
		// {Handler : "CreateNamespace", Method : "POST", Path : "/api/v1/namespaces"},
		// {Handler : "GetNamespaces", Method : "GET", Path : "/api/v1/namespaces"},
		// {Handler : "GetNamespace", Method : "GET", Path : "/api/v1/namespaces/:name"},
		// {Handler : "DeleteNamespace", Method : "DELETE", Path : "/api/v1/namespaces/:name"},

		// //service
		// {Handler : "CreateService", Method : "POST", Path : "/api/v1/namespaces/:namespace/services"},
		// {Handler : "GetService", Method : "GET", Path : "/api/v1/namespaces/:namespace/services/:name"},
		// {Handler : "GetServices", Method : "GET", Path : "/api/v1/namespaces/:namespace/services"},
		// {Handler : "DeleteService", Method : "DELETE", Path : "/api/v1/namespaces/:namespace/services/:name"},

		//configMap
		{Handler : "CreateNamespacedConfigMap", Method : "POST", Path : "/api/v1/namespaces/:namespace/configmaps"},
		{Handler : "ListNamespacedConfigMap", Method : "GET", Path : "/api/v1/namespaces/:namespace/configmaps"},
		{Handler : "DeletecollectionNamespacedConfigMap", Method : "DELETE", Path : "/api/v1/namespaces/:namespace/configmaps"},
		{Handler : "ReadNamespacedConfigMap", Method : "GET", Path : "/api/v1/namespaces/:namespace/configmaps/:name"},
		{Handler : "ReplaceNamespacedConfigMap", Method : "PUT", Path : "/api/v1/namespaces/:namespace/configmaps/:name"},
		{Handler : "DeleteNamespacedConfigMap", Method : "DELETE", Path : "/api/v1/namespaces/:namespace/configmaps/:name"},

		//replicasets
		{Handler : "CreateNamespacedReplicaSet", Method : "POST", Path : "/apis/extensions/v1beta1/namespaces/:namespace/replicasets"},
		{Handler : "ListNamespacedReplicaSet", Method : "GET", Path : "/apis/extensions/v1beta1/namespaces/:namespace/replicasets"},
		{Handler : "DeletecollectionNamespacedReplicaSet", Method : "DELETE", Path : "/apis/extensions/v1beta1/namespaces/:namespace/replicasets"},
		{Handler : "ReadNamespacedReplicaSet", Method : "GET", Path : "/apis/extensions/v1beta1/namespaces/:namespace/replicasets/:name"},
		{Handler : "ReplaceNamespacedReplicaSet", Method : "PUT", Path : "/apis/extensions/v1beta1/namespaces/:namespace/replicasets/:name"},
		{Handler : "DeleteNamespacedReplicaSet", Method : "DELETE", Path : "/apis/extensions/v1beta1/namespaces/:namespace/replicasets/:name"},

		//Secret
		{Handler : "CreateNamespacedSecret", Method : "POST", Path : "/api/v1/namespaces/:namespace/secrets"},
		{Handler : "ListNamespacedSecret", Method : "GET", Path : "/api/v1/namespaces/:namespace/secrets"},
		{Handler : "DeletecollectionNamespacedSecret", Method : "DELETE", Path : "/api/v1/namespaces/:namespace/secrets"},
		{Handler : "ReadNamespacedSecret", Method : "GET", Path : "/api/v1/namespaces/:namespace/secrets/:name"},
		{Handler : "ReplaceNamespacedSecret", Method : "PUT", Path : "/api/v1/namespaces/:namespace/secrets/:name"},
		{Handler : "DeleteNamespacedSecret", Method : "DELETE", Path : "/api/v1/namespaces/:namespace/secrets/:name"},

		//role
		{Handler : "CreateNamespacedRole", Method : "POST", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/:namespace/roles"},
		{Handler : "ListNamespacedRole", Method : "GET", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/:namespace/roles"},
		{Handler : "DeletecollectionNamespacedRole", Method : "DELETE", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/:namespace/roles"},
		{Handler : "ReadNamespacedRole", Method : "GET", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/:namespace/roles/:name"},
		{Handler : "ReplaceNamespacedRole", Method : "PUT", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/:namespace/roles/:name"},
		{Handler : "DeleteNamespacedRole", Method : "DELETE", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/:namespace/roles/:name"},

		//clusterroles
		{Handler : "CreateClusterRole", Method : "POST", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/clusterroles"},
		{Handler : "ListClusterRole", Method : "GET", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/clusterroles"},
		{Handler : "DeletecollectionClusterRole", Method : "DELETE", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/clusterroles"},
		{Handler : "ReadClusterRole", Method : "GET", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/clusterroles/:name"},
		{Handler : "ReplaceClusterRole", Method : "PUT", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/clusterroles/:name"},
		{Handler : "DeleteClusterRole", Method : "DELETE", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/clusterroles/:name"},

		//rolebindings
		{Handler : "CreateNamespacedRoleBinding", Method : "POST", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/:namespace/rolebindings"},
		{Handler : "ListNamespacedRoleBinding", Method : "GET", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/:namespace/rolebindings"},
		{Handler : "DeletecollectionNamespacedRoleBinding", Method : "DELETE", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/:namespace/rolebindings"},
		{Handler : "ReadNamespacedRoleBinding", Method : "GET", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/:namespace/rolebindings/:name"},
		{Handler : "ReplaceNamespacedRoleBinding", Method : "PUT", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/:namespace/rolebindings/:name"},
		{Handler : "DeleteNamespacedRoleBinding", Method : "DELETE", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/:namespace/rolebindings/:name"},

		//clusterrolebindings
		{Handler : "CreateClusterRoleBinding", Method : "POST", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings"},
		{Handler : "ListClusterRoleBinding", Method : "GET", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings"},
		{Handler : "DeletecollectionClusterRoleBinding", Method : "DELETE", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings"},
		{Handler : "ReadClusterRoleBinding", Method : "GET", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings/:name"},
		{Handler : "ReplaceClusterRoleBinding", Method : "PUT", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings/:name"},
		{Handler : "DeleteClusterRoleBinding", Method : "DELETE", Path : "/apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings/:name"},


}