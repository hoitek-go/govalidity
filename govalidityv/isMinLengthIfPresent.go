package govalidityv

func IsMinLengthIfPresent(field string, value interface{}, minLength int) (bool, error) {
	if value == nil {
		return true, nil
	}
	return IsMinLength(field, value, minLength)
}
