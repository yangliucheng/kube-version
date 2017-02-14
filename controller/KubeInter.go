package controller

type KubeInter interface {
	Create()
	Get()
	Put()
	Delete()
}