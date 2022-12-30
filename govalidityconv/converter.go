package govalidityconv

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToNumber(value interface{}) (res *float64, err error) {
	val := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		result := float64(val.Int())
		res = &result
	case uint, uint8, uint16, uint32, uint64:
		result := float64(val.Uint())
		res = &result
	case float32, float64:
		result := float64(val.Float())
		res = &result
	case string:
		result, err := strconv.ParseFloat(val.String(), 0)
		if err != nil {
			res = nil
		} else {
			res = &result
		}
	default:
		err = fmt.Errorf("ToInt: unknown interface type %T", value)
		res = nil
	}

	return
}
