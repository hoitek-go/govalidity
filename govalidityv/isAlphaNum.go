package govalidityv

import (
	"fmt"
	"github.com/hoitek-go/govalidity/govalidityl"
	"github.com/hoitek-go/govalidity/govaliditym"
	"strings"
	"unicode"
)

func IsAlphaNum(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsAlphaNum, label, value)
	str := fmt.Sprintf("%v", value)

	// Get locale
	locale := govalidityl.EnUS
	if len(params) > 1 {
		locale = strings.ToLower(fmt.Sprintf("%v", params[1]))
	}

	// Get unicode range table for the given locale
	prop := &unicode.RangeTable{}
	switch locale {
	case strings.ToLower(govalidityl.FaIR):
		prop = &unicode.RangeTable{
			R16: []unicode.Range16{
				{Lo: 0x0621, Hi: 0x064a, Stride: 1}, // Arabic letters
				{Lo: 0x066e, Hi: 0x06d3, Stride: 1}, // Arabic additional letters
				{Lo: 0x06f0, Hi: 0x06f9, Stride: 1}, // Arabic digits
			},
		}
	}

	// Check if the string contains only alphabetic characters for the given locale
	isValid := true
	for _, r := range str {
		if locale == strings.ToLower(govalidityl.EnUS) || locale == strings.ToLower(govalidityl.FiFI) {
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
				isValid = false
				break
			}
		} else {
			if !unicode.Is(prop, r) {
				isValid = false
				break
			}
		}
	}
	if isValid {
		return true, nil
	}

	return false, err
}
