package govalidityv

import (
	"github.com/hoitek-go/govalidity/govaliditym"
	"regexp"
	"strings"
)

func IsDNSName(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsDNSName, label, value)
	str := value.(string)
	if str == "" || len(strings.Replace(str, ".", "", -1)) > 255 {
		return false, err
	}
	isIp, _ := IsIp(field, str)
	rxDNSName := regexp.MustCompile(`^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`)
	isValid := !isIp && rxDNSName.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}
