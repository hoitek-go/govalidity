package govalidityv

func IsMinIfPresent(field string, value interface{}, min interface{}) (bool, error) {
	if value == nil {
		return true, nil
	}
	return IsMin(field, value, min)
}
