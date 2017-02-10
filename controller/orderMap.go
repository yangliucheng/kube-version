package controller

type OrderMap struct {
	Keys []interface{}
	Segment map[interface{}][]interface{}
}

func NewOrderMap() *OrderMap {

	keys := make([]interface{},0)
	segment := make(map[interface{}][]interface{})

	return &OrderMap {
		Keys: keys,
		Segment: segment,
	}
}

func (orderMap *OrderMap) Set(key interface{}, value ...interface{}) {
	orderMap.Keys = append(orderMap.Keys, key)
	orderMap.Segment[key] = value
}