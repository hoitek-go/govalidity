package govalidityv

import (
	"github.com/hoitek-go/govalidity/govaliditym"
	"time"
)

func IsMaxDate(field string, value interface{}, max time.Time) (bool, error) {
	if value.(time.Time).After(max) {
		return false, GetErrorMessageByFieldValue(govaliditym.Default.IsMaxDate, field, value)
	}
	return true, nil
}
