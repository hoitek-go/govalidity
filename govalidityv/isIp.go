package govalidityv

import (
	"github.com/hoitek-go/govalidity/govaliditym"
	"net"
)

func IsIp(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsIp, label, value)
	str := value.(string)
	isValid := net.ParseIP(str) != nil
	if isValid {
		return true, nil
	}
	return false, err
}
