package fiberBindingHeaderService

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"reflect"
)

func (h *FiberBindingHeader) BindFiberHeader(obj interface{}, c *fiber.Ctx) error {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return errors.New("object must be pointer")
	}
	objType := h.getTypeOfObject(obj)
	objValue := h.getValueOfObject(obj)
	for i := 0; i < objValue.NumField(); i++ {
		if objType.Field(i).Tag.Get(h.TagName) != "" {
			tagVValue := objType.Field(i).Tag.Get(h.TagName)
			objValue.Field(i).Set(reflect.ValueOf(c.Get(tagVValue)))

		}

	}
	return nil
}
