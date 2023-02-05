package govalidityv

import (
	"encoding/json"
	"github.com/hoitek-go/govalidity/govaliditym"
)

func IsJson(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsJson, label, value)
	str := value.(string)
	var js json.RawMessage
	isValid := json.Unmarshal([]byte(str), &js) == nil
	if isValid {
		return true, nil
	}
	return false, err
}
