package govalidityv

import (
	"errors"
	"fmt"
	"github.com/hoitek-go/govalidity/govalidityconv"
	"github.com/hoitek-go/govalidity/govaliditym"
	"github.com/hoitek-go/govalidity/govalidityo"
	"strings"
)

func FilterOperators(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	constrains := params[1].([]string)

	if len(constrains) == 0 {
		constrains = append(constrains,
			govalidityo.EQUALS,
			govalidityo.CONTAINS,
			govalidityo.IS_EMPTY,
			govalidityo.IS_NOT_EMPTY,
			govalidityo.IS_ANY_OF,
			govalidityo.STARTS_WITH,
			govalidityo.ENDS_WITH,
			govalidityo.NUMBER_EQUALS,
			govalidityo.NUMBER_GREATER_THAN,
			govalidityo.NUMBER_GREATER_THAN_EQUALS,
			govalidityo.NUMBER_LESS_THAN,
			govalidityo.NUMBER_LESS_THAN_EQUALS,
			govalidityo.NUMBER_NOT_EQUALS,
		)
	}

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

	errMessage := strings.ReplaceAll(govaliditym.Default.IsFilterOperators, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{in}", fmt.Sprintf("%v", constrains))
	err := errors.New(errMessage)

	found := false
	for _, inVal := range constrains {
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
