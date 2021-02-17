package fiberBindingHeaderService

import (
	"errors"
	"fiberBindingHeader/fiberBindingHeaderService/consts"
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
		tagValue := objType.Field(i).Tag.Get(h.TagName)
		if tagValue != "" {
			headerValue := c.Get(tagValue)

			if headerValue == "" {
				if objType.Field(i).Tag.Get(consts.BindingTag) == consts.RequiredValue {
					return errors.New("fiberBindingHeaderError : " + objType.Field(i).Name + " header required")

				}

			} else {
				fieldValue, err := h.mapValueToStructFieldType(headerValue, objType.Field(i).Type)
				if err != nil {
					return err
				}
				if objValue.Field(i).CanSet() {
					objValue.Field(i).Set(fieldValue)
				}

			}

		}

	}
	return nil
}
