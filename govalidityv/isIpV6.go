package govalidityv

import (
	"github.com/hoitek-go/govalidity/govaliditym"
	"net"
	"strings"
)

func IsIpV6(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsIpV6, label, value)
	str := value.(string)
	ip := net.ParseIP(str)
	isValid := ip != nil && strings.Contains(str, ":")
	if isValid {
		return true, nil
	}
	return false, err
}
