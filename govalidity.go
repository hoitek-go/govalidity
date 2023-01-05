package govalidity

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/hoitek-go/govalidity/govaliditybody"
	"github.com/hoitek-go/govalidity/govalidityconv"
	"github.com/hoitek-go/govalidity/govaliditym"
	"github.com/hoitek-go/govalidity/govalidityv"
)

type (
	Schema                 = map[string]interface{}
	ValidationErrors       = map[string][]error
	ValidityResponseErrors = map[string][]string
	Body                   = map[string]interface{}
	Queries                = map[string]interface{}
	Params                 = map[string]string
)

type FuncSchema struct {
	Fn func(string, ...interface{}) (bool, error)
}

type Validator struct {
	Field        string
	Value        interface{}
	DefaultValue string
	Errors       []error
	Validations  []FuncSchema
}

func New(label string) *Validator {
	return &Validator{
		Field:       label,
		Validations: []FuncSchema{},
	}
}

func (v *Validator) Email() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsEmail,
	})
	return v
}

func (v *Validator) Required() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsRequired,
	})
	return v
}

func (v *Validator) Number() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsNumber,
	})
	return v
}

func (v *Validator) Url() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsUrl,
	})
	return v
}

func (v *Validator) Alpha() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsAlpha,
	})
	return v
}

func (v *Validator) LowerCase() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsLowerCase,
	})
	return v
}

func (v *Validator) UpperCase() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsUpperCase,
	})
	return v
}

func (v *Validator) Int() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsInt,
	})
	return v
}

func (v *Validator) Float() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsFloat,
	})
	return v
}

func (v *Validator) Json() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsJson,
	})
	return v
}

func (v *Validator) Ip() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsIp,
	})
	return v
}

func (v *Validator) IpV4() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsIpV4,
	})
	return v
}

func (v *Validator) IpV6() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsIpV6,
	})
	return v
}

func (v *Validator) Port() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsPort,
	})
	return v
}

func (v *Validator) IsDNSName() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsDNSName,
	})
	return v
}

func (v *Validator) Host() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsHost,
	})
	return v
}

func (v *Validator) Latitude() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsLatitude,
	})
	return v
}

func (v *Validator) Logitude() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsLogitude,
	})
	return v
}

func (v *Validator) AlphaNum() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: govalidityv.IsAlphaNum,
	})
	return v
}

func (v *Validator) InRange(from, to int) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return govalidityv.IsInRange(f, v.Value, from, to)
		},
	})
	return v
}

func (v *Validator) MinMaxLength(min, max int) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return govalidityv.IsMinMaxLength(f, v.Value, min, max)
		},
	})
	return v
}

func (v *Validator) MinLength(min int) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return govalidityv.IsMinLength(f, v.Value, min)
		},
	})
	return v
}

func (v *Validator) MaxLength(max int) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return govalidityv.IsMaxLength(f, v.Value, max)
		},
	})
	return v
}

func (v *Validator) In(in []string) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return govalidityv.IsIn(f, v.Value, in)
		},
	})
	return v
}

func (v *Validator) FilterOperators(operators ...string) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return govalidityv.FilterOperators(f, v.Value, operators)
		},
	})
	return v
}

func (v *Validator) CustomValidator(fn func(string, ...interface{}) (bool, error)) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			isValid, err := fn(v.Field, v.Value)
			label, value := govalidityv.GetFieldLabelAndValue(v.Field, []interface{}{v.Value})
			if err != nil {
				err = govalidityv.GetErrorMessageByFieldValue(err.Error(), label, value)
			}
			return isValid, err
		},
	})
	return v
}

func (v *Validator) Default(value string) *Validator {
	v.DefaultValue = value
	return v
}

func isJson(s string) bool {
	var j map[string]interface{}
	if err := json.Unmarshal([]byte(s), &j); err != nil {
		return false
	}
	return true
}

func convertToMap(s string) map[string]interface{} {
	var j map[string]interface{}
	if err := json.Unmarshal([]byte(s), &j); err != nil {
		return nil
	}
	return j
}

func convertToJson(i interface{}) string {
	bytes, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func sanitizeDataMapToJson(dataMap map[string]interface{}) map[string]interface{} {
	for k, v := range dataMap {
		switch v.(type) {
		case string:
			if isJson(v.(string)) {
				dataMap[k] = convertToMap(v.(string))
				sanitizeDataMapToJson(dataMap[k].(map[string]interface{}))
			}
		}
	}
	return dataMap
}

var validationErrors = ValidationErrors{}
var isValid = true

func validateByJson(baseDataMap map[string]interface{}, dataMap map[string]interface{}, validations Schema, structData interface{}) (bool, ValidationErrors) {

	for k, v := range validations {
		value, ok := dataMap[k]
		if !ok {
			switch v.(type) {
			case *Validator:
				value = ""
				if v.(*Validator).DefaultValue != "" {
					value = v.(*Validator).DefaultValue
					dataMap[k] = value
				}
			case Schema:
				value = Schema{}
			}
		}
		switch v.(type) {
		case Schema:
			valueStr := fmt.Sprintf("%s", value)
			temp := map[string]interface{}{}
			if isJson(valueStr) {
				jsonData := convertToMap(valueStr)
				jsn, ok := jsonData[k]
				if ok {
					validateByJson(baseDataMap, jsn.(Schema), v.(Schema), &temp)
				}
			} else {
				validateByJson(baseDataMap, value.(Schema), v.(Schema), &temp)
			}
		case *Validator:
			v.(*Validator).Value = value
			errs := v.(*Validator).run()
			if len(errs) > 0 {
				isValid = false
				validationErrors[v.(*Validator).Field] = errs
			}
		}
	}
	baseDataMap = sanitizeDataMapToJson(baseDataMap)
	bytes, err := json.Marshal(baseDataMap)
	if err != nil {
		return false, ValidationErrors{
			"error": {
				errors.New("Input Data is Invalid"),
			},
		}
	}
	err = json.Unmarshal(bytes, &structData)
	if err != nil {
		return false, ValidationErrors{
			"error": {
				errors.New("Input Data is Invalid"),
			},
		}
	}
	return isValid, validationErrors
}

func ValidateQueries(r *http.Request, validations Schema, structData interface{}) (bool, ValidationErrors) {
	validationErrors = ValidationErrors{}
	isValid = true
	baseDataMap := Queries{}
	dataMap := Queries{}
	queries := r.URL.Query()
	for k, v := range queries {
		if len(v) > 0 {
			baseDataMap[k] = v[0]
		} else {
			baseDataMap[k] = ""
		}
	}
	for k, v := range baseDataMap {
		value := v.(string)
		if isJson(value) {
			dataMap[k] = `{"` + k + `":` + value + `}`
		} else {
			dataMap[k] = value
		}
	}
	return validateByJson(baseDataMap, dataMap, validations, structData)
}

func ValidateBody(r *http.Request, validations Schema, structData interface{}) (bool, ValidationErrors) {
	validationErrors = ValidationErrors{}
	isValid = true
	dataMap := Body{}
	var baseDataMap Body
	err := govaliditybody.Bind(r, &baseDataMap)
	if err != nil {
		return false, ValidationErrors{
			"UnknownError": []error{
				err,
			},
		}
	}
	for k, v := range baseDataMap {
		switch v.(type) {
		case string:
			if isJson(v.(string)) {
				dataMap[k] = `{"` + k + `":` + v.(string) + `}`
			} else {
				dataMap[k] = v.(string)
			}
		}
	}

	return validateByJson(baseDataMap, dataMap, validations, structData)
}

func ValidateParams(params Params, validations Schema, structData interface{}) (bool, ValidationErrors) {
	validationErrors = ValidationErrors{}
	isValid = true
	baseDataMap := Body{}
	dataMap := Body{}

	for k, v := range params {
		baseDataMap[k] = v
	}

	for k, v := range baseDataMap {
		value := v.(string)
		if isJson(value) {
			dataMap[k] = `{"` + k + `":` + value + `}`
		} else {
			dataMap[k] = value
		}
	}

	return validateByJson(baseDataMap, dataMap, validations, structData)
}

func SetDefaultErrorMessages(v *govaliditym.Validations) {
	govaliditym.SetMessages(v)
}

func SetFieldLabels(l *govaliditym.Labels) {
	govaliditym.SetFieldLables(l)
}

func (v *Validator) run() []error {
	errs := []error{}
	for _, validation := range v.Validations {
		number, errConv := govalidityconv.ToNumber(v.Value)
		str := ""
		if errConv == nil && number != nil {
			str = fmt.Sprintf("%v", *number)
		} else {
			str = v.Value.(string)
		}
		isValid, err := validation.Fn(v.Field, str)
		if !isValid {
			errs = append(errs, err)
		}
	}
	return errs
}

func DumpErrors(errs ValidationErrors) ValidityResponseErrors {
	errMap := map[string][]string{}
	for k, vErrs := range errs {
		for _, e := range vErrs {
			errMap[k] = append(errMap[k], e.Error())
		}
	}
	return errMap
}
