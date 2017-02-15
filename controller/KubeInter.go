package controller

type KubeInter interface {
	Create(out bool)
	Get(out bool)
	Put(out bool)
	Delete(out bool)
}