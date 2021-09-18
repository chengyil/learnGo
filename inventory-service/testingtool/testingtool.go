package testingtool

import "reflect"

type Pair struct {
	funcPtr    interface{}
	originType reflect.Type
	originFunc interface{}
}

func Pairs(fn interface{}) Pair {
	return Pair{
		originType: reflect.ValueOf(fn).Elem().Type(),
		originFunc: reflect.ValueOf(fn).Elem().Interface(),
		funcPtr:    fn,
	}
}

func Restore(p Pair) {
	if p.originFunc == nil {
		reflect.ValueOf(p.funcPtr).Elem().Set(reflect.Zero(p.originType))
	} else {
		reflect.ValueOf(p.funcPtr).Elem().Set(reflect.ValueOf(p.originFunc))
	}
}
