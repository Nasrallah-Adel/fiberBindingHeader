package fiberBindingHeaderService

import (
	"fiberBindingHeader/fiberBindingHeaderService/consts"
)

type FiberBindingHeader struct {
	TagName string
}

func InitFiberBindingHeaderWithTag(tagName string) FiberBindingHeaderInterface {
	if tagName == "" {
		tagName = consts.HeaderDefaultTag
	}
	return &FiberBindingHeader{TagName: tagName}
}
func InitFiberBindingHeader() FiberBindingHeaderInterface {

	return &FiberBindingHeader{TagName: consts.HeaderDefaultTag}
}
