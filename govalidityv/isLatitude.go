package govalidityv

import (
	"github.com/hoitek-go/govalidity/govaliditym"
	"regexp"
)

func IsLatitude(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsLatitude, label, value)
	str := value.(string)
	rxLatitude := regexp.MustCompile("^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$")
	isValid := rxLatitude.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}
