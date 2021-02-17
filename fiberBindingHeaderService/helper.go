package fiberBindingHeaderService

import "reflect"

func (h *FiberBindingHeader) getTypeOfObject(obj interface{}) reflect.Type {
	return reflect.TypeOf(obj).Elem()
}

func (h *FiberBindingHeader) getValueOfObject(obj interface{}) reflect.Value {
	return reflect.ValueOf(obj).Elem()
}
