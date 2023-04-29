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
			govalidityo.NOT_EQUALS,
			govalidityo.GREATER_THAN,
			govalidityo.GREATER_THAN_EQUALS,
			govalidityo.LESS_THAN,
			govalidityo.LESS_THAN_EQUAL,
			govalidityo.LIKE,
			govalidityo.NOT_LIKE,
			govalidityo.IGNORE_LIKE,
			govalidityo.NOT_IGNORE_LIKE,
			govalidityo.IN,
			govalidityo.NOT_IN,
			govalidityo.IS,
			govalidityo.IS_NOT,
			govalidityo.BETWEEN,
			govalidityo.NOT_BETWEEN,
			govalidityo.OVERLAP,
			govalidityo.CONTAINS,
			govalidityo.CONTAINED,
			govalidityo.IS_EMPTY,
			govalidityo.IS_NOT_EMPTY,
			govalidityo.IS_ANY_OF,
			govalidityo.STARTS_WITH,
			govalidityo.ENDS_WITH,
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
