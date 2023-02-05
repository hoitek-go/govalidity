package govalidityv

import (
	"errors"
	"fmt"
	"github.com/hoitek-go/govalidity/govaliditym"
	"strconv"
	"strings"
)

func IsMin(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	min := params[1].(int)

	label, ok := (*govaliditym.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}

	errMessage := govaliditym.Default.IsMin
	errMessage = strings.ReplaceAll(errMessage, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", fmt.Sprintf("%v", dataValue))
	errMessage = strings.ReplaceAll(errMessage, "{min}", strconv.Itoa(min))
	err := errors.New(errMessage)

	if dataValue.(float64) < float64(min) {
		return false, err
	}

	return true, nil
}
