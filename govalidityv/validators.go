package govalidityv

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/hoitek-go/govalidity/govalidityconv"
	"github.com/hoitek-go/govalidity/govaliditym"
)

func GetFieldLabelAndValue(field string, params []interface{}) (string, interface{}) {
	fieldLabel := field
	value := params[0]
	label, ok := (*govaliditym.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}
	return fieldLabel, value
}

func GetErrorMessageByFieldValue(baseErrorMessage string, field string, value interface{}) error {
	errMessage := strings.ReplaceAll(baseErrorMessage, "{field}", field)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value.(string))
	return errors.New(errMessage)
}

func IsEmail(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsEmail, label, value)
	str := value.(string)
	rxEmail := regexp.MustCompile("^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")
	isValid := rxEmail.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}

func IsRequired(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsRequired, label, value)
	str := value.(string)
	if strings.Trim(str, " ") == "" {
		return false, err
	}
	return true, nil
}

func IsNumber(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsNumber, label, value)
	number, numberError := govalidityconv.ToNumber(value)
	if numberError != nil || number == nil {
		return false, err
	}
	return true, nil
}

func IsUrl(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsUrl, label, value)
	str := value.(string)
	if str == "" || utf8.RuneCountInString(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false, err
	}
	strTemp := str
	if strings.Contains(str, ":") && !strings.Contains(str, "://") {
		strTemp = "http://" + str
	}
	u, parseError := url.Parse(strTemp)
	if parseError != nil {
		return false, err
	}
	if strings.HasPrefix(u.Host, ".") {
		return false, err
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false, err
	}
	var (
		IP           string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
		URLSchema    string = `((ftp|tcp|udp|wss?|https?):\/\/)`
		URLUsername  string = `(\S+(:\S*)?@)`
		URLPath      string = `((\/|\?|#)[^\s]*)`
		URLPort      string = `(:(\d{1,5}))`
		URLIP        string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3]|24\d|25[0-5])(\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-5]))`
		URLSubdomain string = `((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))`
	)
	rxURL := regexp.MustCompile(`^` + URLSchema + `?` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?` + URLPath + `?$`)
	isValid := rxURL.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}

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

func IsLowerCase(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsLowerCase, label, value)
	str := value.(string)
	isValid := str == strings.ToLower(str)
	if isValid {
		return true, nil
	}
	return false, err
}

func IsUpperCase(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsUpperCase, label, value)
	str := value.(string)
	isValid := str == strings.ToUpper(str)
	if isValid {
		return true, nil
	}
	return false, err
}

func IsInt(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsInt, label, value)
	str := value.(string)
	rxInt := regexp.MustCompile("^(?:[-+]?(?:0|[1-9][0-9]*))$")
	isValid := rxInt.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}

func IsFloat(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsFloat, label, value)
	str := value.(string)
	rxInt := regexp.MustCompile("^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$")
	isValid := rxInt.MatchString(str)
	if isValid {
		return true, nil
	}
	return false, err
}

func IsJson(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsJson, label, value)
	str := value.(string)
	var js json.RawMessage
	isValid := json.Unmarshal([]byte(str), &js) == nil
	if isValid {
		return true, nil
	}
	return false, err
}

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

func IsIpV4(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsIpV4, label, value)
	str := value.(string)
	ip := net.ParseIP(str)
	isValid := ip != nil && strings.Contains(str, ".")
	if isValid {
		return true, nil
	}
	return false, err
}

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

func IsPort(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsPort, label, value)
	str := value.(string)
	if i, err := strconv.Atoi(str); err == nil && i > 0 && i < 65536 {
		return true, nil
	}
	return false, err
}

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

func IsInRange(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	from := params[1].(int)
	to := params[2].(int)

	label, ok := (*govaliditym.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}

	value := ""
	number, errConv := govalidityconv.ToNumber(dataValue)
	if errConv == nil && number != nil {
		value = fmt.Sprintf("%v", *number)
	}

	errMessage := strings.ReplaceAll(govaliditym.Default.IsInRange, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{from}", strconv.Itoa(from))
	errMessage = strings.ReplaceAll(errMessage, "{to}", strconv.Itoa(to))
	err := errors.New(errMessage)

	if number == nil {
		return false, err
	}

	if int(*number) < from || int(*number) > to {
		return false, err
	}

	return true, nil
}

func IsMinMaxLength(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	min := params[1].(int)
	max := params[2].(int)

	label, ok := (*govaliditym.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}

	value := ""
	number, errConv := govalidityconv.ToNumber(dataValue)
	if errConv == nil && number != nil {
		value = fmt.Sprintf("%v", *number)
	} else {
		value = dataValue.(string)
	}

	errMessage := strings.ReplaceAll(govaliditym.Default.IsMinMaxLength, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{min}", strconv.Itoa(min))
	errMessage = strings.ReplaceAll(errMessage, "{max}", strconv.Itoa(max))
	err := errors.New(errMessage)

	if len(value) < min || len(value) > max {
		return false, err
	}

	return true, nil
}

func IsMinLength(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	min := params[1].(int)

	label, ok := (*govaliditym.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}

	value := ""
	number, errConv := govalidityconv.ToNumber(dataValue)
	if errConv == nil && number != nil {
		value = fmt.Sprintf("%v", *number)
	} else {
		value = dataValue.(string)
	}

	errMessage := strings.ReplaceAll(govaliditym.Default.IsMinLength, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{min}", strconv.Itoa(min))
	err := errors.New(errMessage)

	if len(value) < min {
		return false, err
	}

	return true, nil
}

func IsMaxLength(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	max := params[1].(int)

	label, ok := (*govaliditym.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}

	value := ""
	number, errConv := govalidityconv.ToNumber(dataValue)
	if errConv == nil && number != nil {
		value = fmt.Sprintf("%v", *number)
	} else {
		value = dataValue.(string)
	}

	errMessage := strings.ReplaceAll(govaliditym.Default.IsMaxLength, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{max}", strconv.Itoa(max))
	err := errors.New(errMessage)

	if len(value) > max {
		return false, err
	}

	return true, nil
}

func IsIn(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	in := params[1].([]string)

	label, ok := (*govaliditym.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}

	value := ""
	number, errConv := govalidityconv.ToNumber(dataValue)
	if errConv == nil && number != nil {
		value = fmt.Sprintf("%v", *number)
	} else {
		value = dataValue.(string)
	}

	errMessage := strings.ReplaceAll(govaliditym.Default.IsIn, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{in}", fmt.Sprintf("%v", in))
	err := errors.New(errMessage)

	found := false
	for _, inVal := range in {
		if value == inVal {
			found = true
			break
		}
	}
	if !found {
		return false, err
	}

	return true, nil
}
