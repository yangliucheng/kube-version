package controller

type KubeInter interface {
	Create(handler string)
	Get(handler string)
	Delete(handler string)
}