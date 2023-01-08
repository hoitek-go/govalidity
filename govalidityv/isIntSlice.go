package govalidityv

import (
	"encoding/json"
	"github.com/hoitek-go/govalidity/govaliditym"
)

func IsIntSlice(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsSlice, label, value)
	str, ok := value.(string)
	if !ok {
		return false, err
	}
	sliceData := []interface{}{}
	if err := json.Unmarshal([]byte(str), &sliceData); err != nil {
		return false, err
	}
	isValid := true
	for _, sliceItem := range sliceData {
		switch sliceItem.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		default:
			isValid = false
			break
		}
	}
	if isValid {
		return true, nil
	}
	return false, err
}
