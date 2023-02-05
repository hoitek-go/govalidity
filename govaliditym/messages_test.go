package govaliditym

import (
	"testing"
)

func TestSetMessages(t *testing.T) {
	message := "{field} test {value}"
	SetMessages(&Validations{
		IsEmail: message,
	})
	if (*Default).IsEmail != message {
		t.Error("new error messages are not applied correctly")
	}
}

func TestSetFieldLables(t *testing.T) {
	newLabel := "first name"
	SetFieldLables(&Labels{
		"name": newLabel,
	})
	val, ok := (*FieldLabels)["name"]
	if !ok || val != newLabel {
		t.Error("new field label is not applied correctly")
	}
}
