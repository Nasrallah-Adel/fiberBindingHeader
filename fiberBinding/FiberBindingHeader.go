package fiberBinding

import (
	"errors"
	"fiberBindingHeader/fiberBinding/consts"
	"github.com/gofiber/fiber/v2"
	"reflect"
)

func BindHeader(obj interface{}, c *fiber.Ctx) error {
	TagName := consts.HeaderDefaultTag
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return errors.New("object must be pointer")
	}
	objType := getTypeOfObject(obj)
	objValue := getValueOfObject(obj)
	for i := 0; i < objValue.NumField(); i++ {
		tagValue := objType.Field(i).Tag.Get(TagName)
		if tagValue != "" {
			headerValue := c.Get(tagValue)

			if headerValue == "" {
				if objType.Field(i).Tag.Get(consts.BindingTag) == consts.RequiredValue {
					return errors.New("fiberBindingHeaderError : " + objType.Field(i).Name + " header required")

				}

			} else {
				fieldValue, err := mapValueToStructFieldType(headerValue, objType.Field(i).Type)
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
