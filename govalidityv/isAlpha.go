package govalidityv

import (
	"github.com/hoitek-go/govalidity/govaliditym"
	"regexp"
)

func IsAlpha(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsAlpha, label, value)
	str := value.(string)
	rxAlpha := regexp.MustCompile("^[a-zA-Z]+$")
	isValid := rxAlpha.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}
