package govalidityv

import (
	"github.com/hoitek-go/govalidity/govaliditym"
	"time"
)

func IsMinDate(field string, value interface{}, min time.Time) (bool, error) {
	if value.(time.Time).Before(min) {
		return false, GetErrorMessageByFieldValue(govaliditym.Default.IsMinDate, field, value)
	}
	return true, nil
}
