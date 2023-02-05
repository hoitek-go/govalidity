package govalidityv

import (
	"github.com/hoitek-go/govalidity/govaliditym"
	"time"
)

func IsMinTime(field string, value interface{}, min time.Time) (bool, error) {
	if value.(time.Time).Before(min) {
		return false, GetErrorMessageByFieldValue(govaliditym.Default.IsMinTime, field, value)
	}
	return true, nil
}
