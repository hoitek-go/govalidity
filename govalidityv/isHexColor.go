package govalidityv

import (
	"fmt"
	"github.com/hoitek-go/govalidity/govaliditym"
	"regexp"
)

func IsHexColor(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsHexColor, label, value)
	str := fmt.Sprintf("%v", value)
	rxHexColor := regexp.MustCompile(`^#([a-fA-F0-9]{3}|[a-fA-F0-9]{6})$`)
	isValid := rxHexColor.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}
