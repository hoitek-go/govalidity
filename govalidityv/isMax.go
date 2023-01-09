package govalidityv

import (
	"errors"
	"github.com/hoitek-go/govalidity/govaliditym"
	"strconv"
	"strings"
)

func IsMax(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	max := params[1].(int)

	label, ok := (*govaliditym.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}

	errMessage := govaliditym.Default.IsMax
	errMessage = strings.ReplaceAll(errMessage, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", strconv.Itoa(dataValue.(int)))
	errMessage = strings.ReplaceAll(errMessage, "{max}", strconv.Itoa(max))
	err := errors.New(errMessage)

	if dataValue.(int) > max {
		return false, err
	}

	return true, nil
}
