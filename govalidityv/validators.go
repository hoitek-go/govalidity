package govalidityv

import (
	"errors"
	"github.com/hoitek-go/govalidity/govaliditym"
	"strings"
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
