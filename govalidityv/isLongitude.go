package govalidityv

import (
	"github.com/hoitek-go/govalidity/govaliditym"
	"regexp"
)

func IsLogitude(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsLogitude, label, value)
	str := value.(string)
	rxLogitude := regexp.MustCompile("^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$")
	isValid := rxLogitude.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}
