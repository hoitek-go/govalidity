package govalidityv

import (
	"errors"
	"fmt"
	"github.com/hoitek-go/govalidity/govalidityconv"
	"github.com/hoitek-go/govalidity/govaliditym"
	"strconv"
	"strings"
)

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
