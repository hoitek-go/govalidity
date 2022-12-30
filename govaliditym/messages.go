package govaliditym

import "encoding/json"

var DefaultErrorMessages map[string]string = map[string]string{
	"IsEmail":        "{field} should be a valid email",
	"IsRequired":     "{field} can not be empty",
	"IsNumber":       "{field} can be a valid number",
	"IsUrl":          "{field} can be a valid url",
	"IsAlpha":        "{field} can only contain alphabet characters",
	"IsLowerCase":    "{field} can only be lowercase",
	"IsUpperCase":    "{field} can only be uppercase",
	"IsInt":          "{field} can only be a valid integer number",
	"IsFloat":        "{field} can only be a float number",
	"IsJson":         "{field} can only be json",
	"IsIp":           "{field} can only be a valid ip",
	"IsIpV4":         "{field} can only be a valid ipv4",
	"IsIpV6":         "{field} can only be a valid ipv6",
	"IsPort":         "{field} can only be a valid port",
	"IsDNSName":      "{field} can only be a valid dns name",
	"IsHost":         "{field} can only be a valid host",
	"IsLatitude":     "{field} can only be a valid latitude number",
	"IsLogitude":     "{field} can only be a valid logitude number",
	"IsAlphaNum":     "{field} can only contain alphabet characters and numbers",
	"IsInRange":      "{field} can only be in range {from} and {to}",
	"IsMinMaxLength": "{field} can only have {min} to {max} characters",
	"IsMinLength":    "{field} can only have more than {min} characters",
	"IsMaxLength":    "{field} can only have less than {max} characters",
	"IsIn":           "{field} can only be {in}",
}

type Validations struct {
	IsEmail        string
	IsRequired     string
	IsNumber       string
	IsUrl          string
	IsAlpha        string
	IsLowerCase    string
	IsUpperCase    string
	IsInt          string
	IsFloat        string
	IsJson         string
	IsIp           string
	IsIpV4         string
	IsIpV6         string
	IsPort         string
	IsDNSName      string
	IsHost         string
	IsLatitude     string
	IsLogitude     string
	IsAlphaNum     string
	IsInRange      string
	IsMinMaxLength string
	IsMinLength    string
	IsMaxLength    string
	IsIn           string
}

type Labels = map[string]string

var Default *Validations = &Validations{
	IsEmail:        DefaultErrorMessages["IsEmail"],
	IsRequired:     DefaultErrorMessages["IsRequired"],
	IsNumber:       DefaultErrorMessages["IsNumber"],
	IsUrl:          DefaultErrorMessages["IsUrl"],
	IsAlpha:        DefaultErrorMessages["IsAlpha"],
	IsLowerCase:    DefaultErrorMessages["IsLowerCase"],
	IsUpperCase:    DefaultErrorMessages["IsUpperCase"],
	IsInt:          DefaultErrorMessages["IsInt"],
	IsFloat:        DefaultErrorMessages["IsFloat"],
	IsJson:         DefaultErrorMessages["IsJson"],
	IsIp:           DefaultErrorMessages["IsIp"],
	IsIpV4:         DefaultErrorMessages["IsIpV4"],
	IsIpV6:         DefaultErrorMessages["IsIpV6"],
	IsPort:         DefaultErrorMessages["IsPort"],
	IsDNSName:      DefaultErrorMessages["IsDNSName"],
	IsHost:         DefaultErrorMessages["IsHost"],
	IsLatitude:     DefaultErrorMessages["IsLatitude"],
	IsLogitude:     DefaultErrorMessages["IsLogitude"],
	IsAlphaNum:     DefaultErrorMessages["IsAlphaNum"],
	IsInRange:      DefaultErrorMessages["IsInRange"],
	IsMinMaxLength: DefaultErrorMessages["IsMinMaxLength"],
	IsMinLength:    DefaultErrorMessages["IsMinLength"],
	IsMaxLength:    DefaultErrorMessages["IsMaxLength"],
	IsIn:           DefaultErrorMessages["IsIn"],
}

var FieldLabels *Labels = &Labels{}

func SetMessages(v *Validations) {
	errorMessages := map[string]string{}
	bytes, _ := json.Marshal(v)
	json.Unmarshal(bytes, &errorMessages)
	for key, value := range errorMessages {
		if value == "" {
			errorMessages[key] = DefaultErrorMessages[key]
		}
	}
	bytes, _ = json.Marshal(errorMessages)
	json.Unmarshal(bytes, Default)
}

func SetFieldLables(l *Labels) {
	FieldLabels = l
}
