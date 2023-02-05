package govalidityv

import (
	"errors"
	"fmt"
	"github.com/hoitek-go/govalidity/govalidityconv"
	"github.com/hoitek-go/govalidity/govaliditym"
	"strings"
)

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
