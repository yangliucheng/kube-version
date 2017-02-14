package controller

func init() {}
	var create 	string = "CreateNamespacedReplicaSet"
	var get 	string = "ReadNamespacedReplicaSet"
	var gets 	string = "ListNamespacedReplicaSet"
	var put 	string = "ReplaceNamespacedReplicaSet"
	var delete 	string = "DeletecollectionNamespacedReplicaSet"
}


type KubeReplicaset struct {
	kubeC *KubeClient
	yaml string
}

func NewKubeConfigMap(cli *KubeClient) *KubeReplicaset {
	reslicaset := `{"apiVersion":"extensions/v1beta1","kind":"ReplicaSet","metadata":{"name":"frontend"},"spec":{"replicas":3,"spec":{"containers":[{"name":"redis","image":"redis:latest","ports":[{"containerPort":80}]}]}}}`
	return &KubeReplicaset {
		kubeC: cli,
		yaml : reslicaset,
	}
}

func (kubeReplicaset *KubeReplicaset) Create() {
	body := strings.NewReader(kubeReplicaset.yaml)
	response, err := kubeReplicaset.kubeC.RequestGen.DoHttpRequest(create, easy_http.Mapstring{"namespace": "default"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeReplicaset.kubeC.PrintExcel(response, create)
}

func (kubeReplicaset *KubeReplicaset) Get() {
	response, err := kubeReplicaset.kubeC.RequestGen.DoHttpRequest(get, easy_http.Mapstring{"namespace": "default","name":"example-config"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeReplicaset.kubeC.PrintExcel(response, get)
	response, err = kubeReplicaset.kubeC.RequestGen.DoHttpRequest(gets, easy_http.Mapstring{"namespace": "default"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeReplicaset.kubeC.PrintExcel(response, gets)
}

func (kubeReplicaset *KubeReplicaset) Put() {
	body := strings.NewReader(kubeReplicaset.yaml)
	response, err := kubeReplicaset.kubeC.RequestGen.DoHttpRequest(put, easy_http.Mapstring{"namespace": "default","name":"example-config"}, body, easy_http.Mapstring{"Content-type": "application/json"}, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeReplicaset.kubeC.PrintExcel(response, put)
}

func (kubeReplicaset *KubeReplicaset) Delete() {
	response, err := kubeReplicaset.kubeC.RequestGen.DoHttpRequest(delete, easy_http.Mapstring{"namespace": "default","name":"example-config"}, nil, nil, "")
	if err != nil {
		fmt.Println("send request fail:",err)
		return
	}
	kubeReplicaset.kubeC.PrintExcel(response, delete)
}