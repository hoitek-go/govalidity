package govalidityv

import (
	"github.com/hoitek-go/govalidity/govaliditym"
	"regexp"
)

func IsFilepath(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsFilepath, label, value)
	str := value.(string)
	rxFilepath := regexp.MustCompile("^(?:[-+]?[0-9]+)?(?:[.][0-9]+)?(?:[a-zA-Z0-9]+)?(?:[.][a-zA-Z0-9]+)?$")
	isValid := rxFilepath.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}
