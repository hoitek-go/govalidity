package govalidityv

import (
	"testing"

	"github.com/hoitek-go/govalidity/govaliditym"
)

func TestGetFieldLabelAndValue(t *testing.T) {
	t.Run("GetFieldLabelAndValue", func(t *testing.T) {
		label, value := GetFieldLabelAndValue("test", []interface{}{"val"})
		if label != "test" || value != "val" {
			t.Error("label and value is not valid")
		}
	})

	t.Run("GetFieldLabelAndValue when there is label", func(t *testing.T) {
		govaliditym.SetFieldLables(&govaliditym.Labels{
			"test": "test2",
		})
		label, value := GetFieldLabelAndValue("test", []interface{}{"val"})
		if label != "test2" || value != "val" {
			t.Error("label and value is not valid")
		}
	})
}

func TestGetErrorMessageByFieldValue(t *testing.T) {
	t.Run("GetErrorMessageByFieldValue", func(t *testing.T) {
		base := "{field} with {value}"
		err := GetErrorMessageByFieldValue(base, "test", "val")
		if err.Error() != "test with val" {
			t.Error("error message is not valid")
		}
	})
}

type ValidationsTest struct {
	Fn            func(field string, params ...interface{}) (bool, error)
	Name          string
	Field         string
	Params        interface{}
	ExpectedValid bool
	ExpectedError error
}

var validationTestCases []ValidationsTest = []ValidationsTest{
	{Fn: IsEmail, Name: "IsEmail", Field: "email", Params: "test@email.com", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsEmail, Name: "IsEmail", Field: "email", Params: "test", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsEmail, "email", "test")},
	{Fn: IsRequired, Name: "IsRequired", Field: "email", Params: "test@email.com", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsRequired, Name: "IsRequired", Field: "email", Params: "", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsRequired, "email", "")},
	{Fn: IsNumber, Name: "IsNumber", Field: "age", Params: "20", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsNumber, Name: "IsNumber", Field: "age", Params: "test", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsNumber, "age", "test")},
	{Fn: IsUrl, Name: "IsUrl", Field: "url", Params: "https://google.com", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsUrl, Name: "IsUrl", Field: "url", Params: "http://google.com", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsUrl, Name: "IsUrl", Field: "url", Params: "https://www.google.com", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsUrl, Name: "IsUrl", Field: "url", Params: "http://www.google.com", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsUrl, Name: "IsUrl", Field: "url", Params: "www.google.com", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsUrl, Name: "IsUrl", Field: "url", Params: "test", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsUrl, "url", "test")},
	{Fn: IsUrl, Name: "IsUrl", Field: "url", Params: "", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsUrl, "url", "")},
	{Fn: IsUrl, Name: "IsUrl", Field: "url", Params: "test:aa", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsUrl, "url", "test:aa")},
	{Fn: IsUrl, Name: "IsUrl", Field: "url", Params: "https://.a.google.com", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsUrl, "url", "https://.a.google.com")},
	{Fn: IsUrl, Name: "IsUrl", Field: "url", Params: "https://google.com.a.b", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsUrl, "url", "https://google.com.a.b")},
	{Fn: IsAlpha, Name: "IsAlpha", Field: "name", Params: "test", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsAlpha, Name: "IsAlpha", Field: "name", Params: "test1", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsAlpha, "name", "test1")},
	{Fn: IsAlpha, Name: "IsAlpha", Field: "name", Params: "234", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsAlpha, "name", "234")},
	{Fn: IsLowerCase, Name: "IsLowerCase", Field: "name", Params: "test", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsLowerCase, Name: "IsLowerCase", Field: "name", Params: "tesT", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsLowerCase, "name", "tesT")},
	{Fn: IsLowerCase, Name: "IsLowerCase", Field: "name", Params: "TEST", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsLowerCase, "name", "TEST")},
	{Fn: IsUpperCase, Name: "IsUpperCase", Field: "name", Params: "TEST", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsUpperCase, Name: "IsUpperCase", Field: "name", Params: "tesT", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsUpperCase, "name", "tesT")},
	{Fn: IsUpperCase, Name: "IsUpperCase", Field: "name", Params: "test", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsUpperCase, "name", "test")},
	{Fn: IsInt, Name: "IsInt", Field: "age", Params: "20", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsInt, Name: "IsInt", Field: "age", Params: "test", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsInt, "age", "test")},
	{Fn: IsInt, Name: "IsInt", Field: "age", Params: "1.23", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsInt, "age", "1.23")},
	{Fn: IsFloat, Name: "IsFloat", Field: "age", Params: "20.3", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsFloat, Name: "IsFloat", Field: "age", Params: "test", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsFloat, "age", "test")},
	{Fn: IsFloat, Name: "IsFloat", Field: "age", Params: "1", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsJson, Name: "IsJson", Field: "json", Params: "[{\"key\":\"value\"}]", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsJson, Name: "IsJson", Field: "json", Params: "test", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsJson, "json", "test")},
	{Fn: IsIp, Name: "IsIp", Field: "ip", Params: "127.0.0.1", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsIp, Name: "IsIp", Field: "ip", Params: "127.0.0.", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsIp, "ip", "127.0.0.")},
	{Fn: IsIp, Name: "IsIp", Field: "ip", Params: "127.0.0.2345", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsIp, "ip", "127.0.0.2345")},
	{Fn: IsIp, Name: "IsIp", Field: "ip", Params: "test", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsIp, "ip", "test")},
	{Fn: IsIpV4, Name: "IsIpV4", Field: "ip", Params: "127.0.0.1", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsIpV4, Name: "IsIpV4", Field: "ip", Params: "127.0.0.", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsIpV4, "ip", "127.0.0.")},
	{Fn: IsIpV4, Name: "IsIpV4", Field: "ip", Params: "127.0.0.2345", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsIpV4, "ip", "127.0.0.2345")},
	{Fn: IsIpV4, Name: "IsIpV4", Field: "ip", Params: "2001:db8:3333:4444:5555:6666:7777:8888", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsIpV4, "ip", "2001:db8:3333:4444:5555:6666:7777:8888")},
	{Fn: IsIpV6, Name: "IsIpV6", Field: "ip", Params: "2001:db8:3333:4444:5555:6666:7777:8888", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsIpV6, Name: "IsIpV6", Field: "ip", Params: "127.0.0.", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsIpV6, "ip", "127.0.0.")},
	{Fn: IsIpV6, Name: "IsIpV6", Field: "ip", Params: "127.0.0.2345", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsIpV6, "ip", "127.0.0.2345")},
	{Fn: IsIpV6, Name: "IsIpV6", Field: "ip", Params: "127.0.0.1", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsIpV6, "ip", "127.0.0.1")},
	{Fn: IsPort, Name: "IsPort", Field: "port", Params: "5600", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsPort, Name: "IsPort", Field: "port", Params: "-1", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsPort, "port", "-1")},
	{Fn: IsPort, Name: "IsPort", Field: "port", Params: "65536", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsPort, "port", "65536")},
	{Fn: IsDNSName, Name: "IsDNSName", Field: "dns", Params: "www.google.com", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsDNSName, Name: "IsDNSName", Field: "dns", Params: "127.0.0.1", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsDNSName, "dns", "127.0.0.1")},
	{Fn: IsDNSName, Name: "IsDNSName", Field: "dns", Params: "", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsDNSName, "dns", "")},
	{Fn: IsDNSName, Name: "IsDNSName", Field: "dns", Params: "test", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsHost, Name: "IsHost", Field: "host", Params: "test", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsHost, Name: "IsHost", Field: "host", Params: "www.google.com", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsHost, Name: "IsHost", Field: "host", Params: "127.0.0.1", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsHost, Name: "IsHost", Field: "host", Params: "..test.", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsHost, "host", "..test.")},
	{Fn: IsLatitude, Name: "IsLatitude", Field: "lat", Params: "-15.569", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsLatitude, Name: "IsLatitude", Field: "lat", Params: "-150.569", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsLatitude, "lat", "-150.569")},
	{Fn: IsLatitude, Name: "IsLatitude", Field: "lat", Params: "15.569", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsLatitude, Name: "IsLatitude", Field: "lat", Params: "150.569", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsLatitude, "lat", "150.569")},
	{Fn: IsLogitude, Name: "IsLogitude", Field: "lng", Params: "-77.0364", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsLogitude, Name: "IsLogitude", Field: "lng", Params: "-1500.569", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsLogitude, "lng", "-1500.569")},
	{Fn: IsLogitude, Name: "IsLogitude", Field: "lng", Params: "77.0364", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsLogitude, Name: "IsLogitude", Field: "lng", Params: "1500.569", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsLogitude, "lng", "1500.569")},
	{Fn: IsAlphaNum, Name: "IsAlphaNum", Field: "name", Params: "77", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsAlphaNum, Name: "IsAlphaNum", Field: "name", Params: "test", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsAlphaNum, Name: "IsAlphaNum", Field: "name", Params: "77test", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsAlphaNum, Name: "IsAlphaNum", Field: "name", Params: "test77", ExpectedValid: true, ExpectedError: nil},
	{Fn: IsAlphaNum, Name: "IsAlphaNum", Field: "name", Params: ".s", ExpectedValid: false, ExpectedError: GetErrorMessageByFieldValue(govaliditym.Default.IsAlphaNum, "name", ".s")},
}

func TestValidations(t *testing.T) {

	for _, vTestCase := range validationTestCases {
		isValid, err := vTestCase.Fn(vTestCase.Field, vTestCase.Params)
		if isValid != vTestCase.ExpectedValid {
			t.Errorf("%s: validation result is not the same as expected result", vTestCase.Name)
		}
		if (err == nil && vTestCase.ExpectedError != nil) || (err != nil && vTestCase.ExpectedError == nil) {
			t.Errorf("%s: validation error is not the same as expected error", vTestCase.Name)
		}
		if err != nil && vTestCase.ExpectedError != nil {
			if err.Error() != vTestCase.ExpectedError.Error() {
				t.Errorf("%s: validation error is not the same as expected error", vTestCase.Name)
			}
		}
	}
}

func TestIsInRange(t *testing.T) {
	t.Run("IsInRange not in range", func(t *testing.T) {
		isValid, err := IsInRange("number", "30", 10, 20)
		if err == nil {
			t.Error("should throw error when number is not in range")
		}
		if isValid {
			t.Error("result should not be valid because there is error")
		}
	})
	t.Run("IsInRange", func(t *testing.T) {
		isValid, err := IsInRange("number", "15", 10, 20)
		if err != nil {
			t.Error("should not throw error when number is in range")
		}
		if !isValid {
			t.Error("result should be valid because there is no error")
		}
	})
	t.Run("IsInRange with label", func(t *testing.T) {
		govaliditym.SetFieldLables(&govaliditym.Labels{
			"number": "num",
		})
		isValid, err := IsInRange("number", "15", 10, 20)
		if err != nil {
			t.Error("should not throw error when number is in range")
		}
		if !isValid {
			t.Error("result should be valid because there is no error")
		}
	})
	t.Run("IsInRange with label", func(t *testing.T) {
		govaliditym.SetFieldLables(&govaliditym.Labels{
			"number": "num",
		})
		isValid, err := IsInRange("number", "test", 10, 20)
		if err == nil {
			t.Error("should throw error when string is not number")
		}
		if isValid {
			t.Error("result should not be valid because there is no error")
		}
	})
}

func TestIsMinMaxLength(t *testing.T) {
	t.Run("IsMinMaxLength not in range", func(t *testing.T) {
		isValid, err := IsMinMaxLength("name", "sample", 10, 20)
		if err == nil {
			t.Error("should throw error when string is not in selected min max")
		}
		if isValid {
			t.Error("result should not be valid because there is error")
		}
	})
	t.Run("IsMinMaxLength", func(t *testing.T) {
		isValid, err := IsMinMaxLength("number", "sample", 3, 10)
		if err != nil {
			t.Error("should not throw error when string is in selected min max")
		}
		if !isValid {
			t.Error("result should be valid because there is no error")
		}
	})
}

func TestIsMinLength(t *testing.T) {
	t.Run("IsMinLength not in range", func(t *testing.T) {
		isValid, err := IsMinLength("name", "sample", 10)
		if err == nil {
			t.Error("should throw error when string is not in selected min")
		}
		if isValid {
			t.Error("result should not be valid because there is error")
		}
	})
	t.Run("IsMinLength", func(t *testing.T) {
		isValid, err := IsMinLength("number", "sample", 3)
		if err != nil {
			t.Error("should not throw error when string is in selected min")
		}
		if !isValid {
			t.Error("result should be valid because there is no error")
		}
	})
}

func TestIsMaxLength(t *testing.T) {
	t.Run("IsMaxLength not in range", func(t *testing.T) {
		isValid, err := IsMaxLength("name", "sample", 2)
		if err == nil {
			t.Error("should throw error when string is not in selected max")
		}
		if isValid {
			t.Error("result should not be valid because there is error")
		}
	})
	t.Run("IsMaxLength", func(t *testing.T) {
		isValid, err := IsMaxLength("number", "sample", 10)
		if err != nil {
			t.Error("should not throw error when string is in selected max")
		}
		if !isValid {
			t.Error("result should be valid because there is no error")
		}
	})
}

func TestIsIn(t *testing.T) {
	t.Run("TestIsIn not in range", func(t *testing.T) {
		isValid, err := IsIn("name", "sample", []string{"test"})
		if err == nil {
			t.Error("should throw error when string is not in selected list")
		}
		if isValid {
			t.Error("result should not be valid because there is error")
		}
	})
	t.Run("TestIsIn", func(t *testing.T) {
		isValid, err := IsIn("name", "sample", []string{"test", "sample"})
		if err != nil {
			t.Error("should not throw error when string is in selected list")
		}
		if !isValid {
			t.Error("result should be valid because there is no error")
		}
	})
	t.Run("TestIsIn with label", func(t *testing.T) {
		govaliditym.SetFieldLables(&govaliditym.Labels{
			"name": "first name",
		})
		isValid, err := IsIn("name", "sample", []string{"test", "sample"})
		if err != nil {
			t.Error("should not throw error when string is in selected list")
		}
		if !isValid {
			t.Error("result should be valid because there is no error")
		}
	})
}
