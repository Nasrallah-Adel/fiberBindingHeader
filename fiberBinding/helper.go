package fiberBinding

import (
	"errors"
	"reflect"
	"strconv"
)

func getTypeOfObject(obj interface{}) reflect.Type {
	return reflect.TypeOf(obj).Elem()
}

func getValueOfObject(obj interface{}) reflect.Value {
	return reflect.ValueOf(obj).Elem()
}
func mapValueToStructFieldType(value string, valueType reflect.Type) (reflect.Value, error) {
	var result reflect.Value
	var convertError = errors.New("error parse " + value + " to " + valueType.String())
	switch valueType.Kind() {
	case reflect.Float32:
		f32, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return reflect.Value{}, convertError
		}
		result = reflect.ValueOf(float32(f32))
	case reflect.Float64:
		f64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return reflect.Value{}, convertError
		}
		result = reflect.ValueOf(f64)
	case reflect.String:
		result = reflect.ValueOf(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return reflect.Value{}, convertError
		}
		result = reflect.ValueOf(int(i))
	case reflect.Int8:
		i8, err := strconv.ParseInt(value, 10, 8)
		if err != nil {
			return reflect.Value{}, convertError
		}
		result = reflect.ValueOf(int8(i8))
	case reflect.Int16:
		i16, err := strconv.ParseInt(value, 10, 16)
		if err != nil {
			return reflect.Value{}, convertError
		}
		result = reflect.ValueOf(int16(i16))
	case reflect.Int32:
		i32, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return reflect.Value{}, convertError
		}
		result = reflect.ValueOf(int32(i32))
	case reflect.Int64:
		i64, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return reflect.Value{}, convertError
		}
		result = reflect.ValueOf(i64)
	case reflect.Bool:
		bo, err := strconv.ParseBool(value)
		if err != nil {
			return reflect.Value{}, convertError
		}
		result = reflect.ValueOf(bo)
	case reflect.Uint:
		u, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return reflect.Value{}, convertError
		}
		result = reflect.ValueOf(uint(u))
	case reflect.Uint8:
		u8, err := strconv.ParseUint(value, 10, 8)
		if err != nil {
			return reflect.Value{}, convertError
		}
		result = reflect.ValueOf(uint8(u8))
	case reflect.Uint16:
		u16, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return reflect.Value{}, convertError
		}
		result = reflect.ValueOf(uint16(u16))
	case reflect.Uint32:
		u32, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return reflect.Value{}, convertError
		}
		result = reflect.ValueOf(uint32(u32))
	case reflect.Uint64:
		u64, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return reflect.Value{}, convertError
		}
		result = reflect.ValueOf(u64)

	}
	return result, nil
}
