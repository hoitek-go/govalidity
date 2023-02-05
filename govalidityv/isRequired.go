package govalidityv

import (
	"github.com/hoitek-go/govalidity/govaliditym"
	"strings"
)

func IsRequired(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsRequired, label, value)
	str := value.(string)
	if strings.Trim(str, " ") == "" {
		return false, err
	}
	return true, nil
}
