package govalidityv

func IsMaxLengthIfPresent(field string, value interface{}, maxLength int) (bool, error) {
	if value == nil {
		return true, nil
	}
	return IsMaxLength(field, value, maxLength)
}
