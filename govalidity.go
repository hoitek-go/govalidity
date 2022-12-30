package govalidity

import (
	"encoding/json"
	"net/http"

	"github.com/hoitek-go/govalidity/govaliditybody"
	"github.com/hoitek-go/govalidity/govaliditym"
	"github.com/hoitek-go/govalidity/govalidityv"
)

type Schema = map[string]*Validator
type ValidationErrors = map[string][]error
type Body = map[string]interface{}

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

func New() *Validator {
	return &Validator{
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

func (v *Validator) AlpaNum() *Validator {
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

func ValidateBody(r *http.Request, validations Schema) (bool, ValidationErrors, Body) {
	var dataMap Body
	err := govaliditybody.Bind(r, &dataMap)
	if err != nil {
		return false, ValidationErrors{
			"UnknownError": []error{
				err,
			},
		}, nil
	}
	validationErrors := ValidationErrors{}
	isValid := true
	for k, v := range validations {
		value, ok := dataMap[k]
		if !ok {
			value = ""
			if v.DefaultValue != "" {
				value = v.DefaultValue
				dataMap[k] = value
			}
		}
		v.Field = k
		v.Value = value
		errs := v.run()
		if len(errs) > 0 {
			isValid = false
		}
		if len(errs) > 0 {
			validationErrors[v.Field] = errs
		}
	}
	return isValid, validationErrors, dataMap
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
		isValid, err := validation.Fn(v.Field, v.Value.(string))
		if !isValid {
			errs = append(errs, err)
		}
	}
	return errs
}

func GetBodyFromJson(dataMap Body, result interface{}) error {
	bytes, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, result)
	if err != nil {
		return err
	}
	return nil
}
