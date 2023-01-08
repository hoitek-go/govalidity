package govalidityv

import "github.com/hoitek-go/govalidity/govaliditym"

func IsHost(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsHost, label, value)
	str := value.(string)
	isIp, _ := IsIp(field, str)
	isDNSName, _ := IsDNSName(field, str)
	isValid := isIp || isDNSName
	if isValid {
		return true, nil
	}
	return false, err
}
