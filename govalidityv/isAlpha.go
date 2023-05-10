package govalidityv

import (
	"fmt"
	"github.com/hoitek-go/govalidity/govalidityl"
	"github.com/hoitek-go/govalidity/govaliditym"
	"strings"
	"unicode"
)

func IsAlpha(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(govaliditym.Default.IsAlpha, label, value)
	str := fmt.Sprintf("%v", value)

	// Get locale
	locale := govalidityl.EnUS
	if len(params) > 1 {
		locale = strings.ToLower(fmt.Sprintf("%v", params[1]))
	}

	// Get unicode range table for the given locale
	prop := &unicode.RangeTable{}
	switch locale {
	case strings.ToLower(govalidityl.EnUS):
		prop = unicode.Letter
	case strings.ToLower(govalidityl.FiFI):
		prop = &unicode.RangeTable{
			R16: []unicode.Range16{
				{Lo: 0x0041, Hi: 0x005a, Stride: 1}, // A-Z
				{Lo: 0x0061, Hi: 0x007a, Stride: 1}, // a-z
				{Lo: 0x00c4, Hi: 0x00e4, Stride: 1}, // Ä-ä
				{Lo: 0x00d6, Hi: 0x00f6, Stride: 1}, // Ö-ö
			},
		}
	case strings.ToLower(govalidityl.FaIR):
		prop = &unicode.RangeTable{
			R16: []unicode.Range16{
				{Lo: 0x0621, Hi: 0x064a, Stride: 1}, // Arabic letters
				{Lo: 0x066e, Hi: 0x06d3, Stride: 1}, // Arabic additional letters
			},
		}
	default:
		prop = unicode.Letter
	}

	// Check if the string contains only alphabetic characters for the given locale
	isValid := true
	for _, r := range str {
		if !unicode.Is(prop, r) {
			isValid = false
			break
		}
	}
	if isValid {
		return true, nil
	}

	return false, err
}
