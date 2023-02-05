package govalidityv

import (
	"github.com/hoitek-go/govalidity/govaliditym"
	"strings"
)

func IsUpperCase(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsUpperCase, label, value)
	str := value.(string)
	isValid := str == strings.ToUpper(str)
	if isValid {
		return true, nil
	}
	return false, err
}
