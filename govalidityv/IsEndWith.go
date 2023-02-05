package govalidityv

import (
	"errors"
	"fmt"
	"github.com/hoitek-go/govalidity/govalidityconv"
	"github.com/hoitek-go/govalidity/govaliditym"
	"strings"
)

func IsEndWith(field string, dataValue interface{}, str string) (bool, error) {
	fieldLabel := field

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

	if strings.Index(value, str) == len(value)-1 {
		return true, nil
	}

	errMessage := strings.ReplaceAll(govaliditym.Default.IsMaxLength, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{str}", str)
	err := errors.New(errMessage)

	return false, err
}
