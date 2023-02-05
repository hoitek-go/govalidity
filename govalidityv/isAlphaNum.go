package govalidityv

import (
	"github.com/hoitek-go/govalidity/govaliditym"
	"regexp"
)

func IsAlphaNum(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsAlphaNum, label, value)
	str := value.(string)
	rxAlphaNum := regexp.MustCompile("^[a-zA-Z0-9]+$")
	isValid := rxAlphaNum.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}
