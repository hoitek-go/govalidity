package govalidity

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/hoitek-go/govalidity/govaliditym"
)

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestNew(t *testing.T) {
	v := New()
	if v == nil {
		t.Error("new instance of validator struct can not be nil")
	}
}

func TestEmail(t *testing.T) {
	v := New()
	v.Email()
	if len(v.Validations) <= 0 {
		t.Error("Email validator should be set in validations")
	}
}

func TestRequired(t *testing.T) {
	v := New()
	v.Required()
	if len(v.Validations) <= 0 {
		t.Error("Required validator should be set in validations")
	}
}

func TestNumber(t *testing.T) {
	v := New()
	v.Number()
	if len(v.Validations) <= 0 {
		t.Error("Number validator should be set in validations")
	}
}

func TestUrl(t *testing.T) {
	v := New()
	v.Url()
	if len(v.Validations) <= 0 {
		t.Error("Url validator should be set in validations")
	}
}

func TestAlpha(t *testing.T) {
	v := New()
	v.Alpha()
	if len(v.Validations) <= 0 {
		t.Error("Alpha validator should be set in validations")
	}
}

func TestLowerCase(t *testing.T) {
	v := New()
	v.LowerCase()
	if len(v.Validations) <= 0 {
		t.Error("LowerCase validator should be set in validations")
	}
}

func TestUpperCase(t *testing.T) {
	v := New()
	v.UpperCase()
	if len(v.Validations) <= 0 {
		t.Error("UpperCase validator should be set in validations")
	}
}

func TestInt(t *testing.T) {
	v := New()
	v.Int()
	if len(v.Validations) <= 0 {
		t.Error("Int validator should be set in validations")
	}
}

func TestFloat(t *testing.T) {
	v := New()
	v.Float()
	if len(v.Validations) <= 0 {
		t.Error("Float validator should be set in validations")
	}
}

func TestJson(t *testing.T) {
	v := New()
	v.Json()
	if len(v.Validations) <= 0 {
		t.Error("Json validator should be set in validations")
	}
}

func TestIp(t *testing.T) {
	v := New()
	v.Ip()
	if len(v.Validations) <= 0 {
		t.Error("Ip validator should be set in validations")
	}
}

func TestIpV4(t *testing.T) {
	v := New()
	v.IpV4()
	if len(v.Validations) <= 0 {
		t.Error("IpV4 validator should be set in validations")
	}
}

func TestIpV6(t *testing.T) {
	v := New()
	v.IpV6()
	if len(v.Validations) <= 0 {
		t.Error("IpV6 validator should be set in validations")
	}
}

func TestPort(t *testing.T) {
	v := New()
	v.Port()
	if len(v.Validations) <= 0 {
		t.Error("Port validator should be set in validations")
	}
}

func TestIsDNSName(t *testing.T) {
	v := New()
	v.IsDNSName()
	if len(v.Validations) <= 0 {
		t.Error("IsDNSName validator should be set in validations")
	}
}

func TestHost(t *testing.T) {
	v := New()
	v.Host()
	if len(v.Validations) <= 0 {
		t.Error("Host validator should be set in validations")
	}
}

func TestLatitude(t *testing.T) {
	v := New()
	v.Latitude()
	if len(v.Validations) <= 0 {
		t.Error("Latitude validator should be set in validations")
	}
}

func TestLogitude(t *testing.T) {
	v := New()
	v.Logitude()
	if len(v.Validations) <= 0 {
		t.Error("Logitude validator should be set in validations")
	}
}

func TestAlpaNum(t *testing.T) {
	v := New()
	v.AlpaNum()
	if len(v.Validations) <= 0 {
		t.Error("AlpaNum validator should be set in validations")
	}
}

func TestInRange(t *testing.T) {
	v := New()
	v.InRange(0, 10)
	if len(v.Validations) <= 0 {
		t.Error("InRange validator should be set in validations")
	}
}

func TestMinMaxLength(t *testing.T) {
	v := New()
	v.MinMaxLength(0, 10)
	if len(v.Validations) <= 0 {
		t.Error("MinMaxLength validator should be set in validations")
	}
}

func TestMinLength(t *testing.T) {
	v := New()
	v.MinLength(0)
	if len(v.Validations) <= 0 {
		t.Error("MinLength validator should be set in validations")
	}
}

func TestMaxLength(t *testing.T) {
	v := New()
	v.MaxLength(10)
	if len(v.Validations) <= 0 {
		t.Error("MaxLength validator should be set in validations")
	}
}

func TestIn(t *testing.T) {
	v := New()
	v.In([]string{"test"})
	if len(v.Validations) <= 0 {
		t.Error("In validator should be set in validations")
	}
}

func TestCustomValidator(t *testing.T) {
	v := New()
	v.CustomValidator(func(s string, i ...interface{}) (bool, error) {
		return false, errors.New("test")
	})
	if len(v.Validations) <= 0 {
		t.Error("CustomValidator should be set in validations")
	}
}

func TestDefault(t *testing.T) {
	v := New()
	v.Default("test")
	if v.DefaultValue != "test" {
		t.Error("Default value is not set correctly")
	}
}

func TestValidateBody(t *testing.T) {
	t.Run("BodyIsNil", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		schema := Schema{
			"email":        New().Email().Required(),
			"name":         New().LowerCase().In([]string{"saeed", "taher"}).Required(),
			"age":          New().Number().Required(),
			"url":          New().Url().Required(),
			"json":         New().Json(),
			"ip":           New().Ip().Required(),
			"filter[page]": New().Int().InRange(10, 20).Required(),
		}
		isValid, _, _ := ValidateBody(r, schema)
		if isValid {
			t.Error("Should throw error when body is nil")
		}
	})

	t.Run("BodyIsNotPrepare", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", errReader(0))
		schema := Schema{
			"email":        New().Email().Required(),
			"name":         New().LowerCase().In([]string{"saeed", "taher"}).Required(),
			"age":          New().Number().Default("20").Required(),
			"url":          New().Url().Required(),
			"json":         New().Json(),
			"ip":           New().Ip().Required(),
			"filter[page]": New().Int().InRange(10, 20).Required(),
		}
		isValid, _, _ := ValidateBody(r, schema)
		if isValid {
			t.Error("Should throw error when body is not prepare")
		}
	})

	t.Run("BodyIsPrepare", func(t *testing.T) {
		r := &http.Request{
			Form: url.Values{
				"name": []string{
					"test",
				},
			},
			Body: io.NopCloser(strings.NewReader("{\"email\":\"sgh370@yahoo.com\",\"name\":\"saeed\",\"lastName\":\"ghanbari\",\"age\":\"50\",\"url\":\"https://google.com\",\"json\":\"{\\\"key\\\":\\\"value\\\"}\",\"ip\":\"127.0.0.1\",\"filter[page]\":\"12\"}")),
		}
		schema := Schema{
			"email":    New().Email().MinLength(5).Required(),
			"name":     New().LowerCase().MinMaxLength(3, 20).In([]string{"saeed", "taher"}).Required(),
			"lastName": New().LowerCase().MaxLength(25).Required(),
			"age":      New().Number().Default("20").Required(),
			"url":      New().Url().Required(),
			"json":     New().Json(),
			"ip":       New().Ip().Required(),
			"filter[page]": New().Int().InRange(10, 20).CustomValidator(func(s string, i ...interface{}) (bool, error) {
				return true, nil
			}).Required(),
		}
		isValid, _, _ := ValidateBody(r, schema)
		if !isValid {
			t.Error("should be valid")
		}
	})

	t.Run("CustomValidatorError", func(t *testing.T) {
		r := &http.Request{
			Form: url.Values{
				"name": []string{
					"test",
				},
			},
			Body: io.NopCloser(strings.NewReader("{\"email\":\"sgh370@yahoo.com\",\"name\":\"saeed\",\"lastName\":\"ghanbari\",\"age\":\"50\",\"url\":\"https://google.com\",\"json\":\"{\\\"key\\\":\\\"value\\\"}\",\"ip\":\"127.0.0.1\",\"filter[page]\":\"12\"}")),
		}
		schema := Schema{
			"email":    New().Email().MinLength(5).Required(),
			"name":     New().LowerCase().MinMaxLength(3, 20).In([]string{"saeed", "taher"}).Required(),
			"lastName": New().LowerCase().MaxLength(25).Required(),
			"age":      New().Number().Default("20").Required(),
			"url":      New().Url().Required(),
			"json":     New().Json(),
			"ip":       New().Ip().Required(),
			"filter[page]": New().Int().InRange(10, 20).CustomValidator(func(s string, i ...interface{}) (bool, error) {
				return false, errors.New("test")
			}).Required(),
		}
		isValid, _, _ := ValidateBody(r, schema)
		if isValid {
			t.Error("should be invalid")
		}
	})

	t.Run("Default", func(t *testing.T) {
		r := &http.Request{
			Form: url.Values{
				"name": []string{
					"test",
				},
			},
			Body: io.NopCloser(strings.NewReader("{\"email\":\"sgh370@yahoo.com\",\"name\":\"saeed\",\"lastName\":\"ghanbari\",\"url\":\"https://google.com\",\"json\":\"{\\\"key\\\":\\\"value\\\"}\",\"ip\":\"127.0.0.1\",\"filter[page]\":\"12\"}")),
		}
		schema := Schema{
			"email":    New().Email().MinLength(5).Required(),
			"name":     New().LowerCase().MinMaxLength(3, 20).In([]string{"saeed", "taher"}).Required(),
			"lastName": New().LowerCase().MaxLength(25).Required(),
			"age":      New().Number().Default("20").Required(),
			"url":      New().Url().Required(),
			"json":     New().Json(),
			"ip":       New().Ip().Required(),
			"filter[page]": New().Int().InRange(10, 20).CustomValidator(func(s string, i ...interface{}) (bool, error) {
				return false, errors.New("test")
			}).Required(),
		}
		isValid, _, _ := ValidateBody(r, schema)
		if isValid {
			t.Error("should be invalid")
		}
	})
}

func TestSetDefaultErrorMessages(t *testing.T) {
	SetDefaultErrorMessages(&govaliditym.Validations{
		IsEmail: "test",
	})
	if govaliditym.Default.IsEmail != "test" {
		t.Error("custom error message is not set properly")
	}
}

func TestSetFieldLabels(t *testing.T) {
	SetFieldLabels(&govaliditym.Labels{
		"name": "test",
	})
	name, ok := (*govaliditym.FieldLabels)["name"]
	if !ok || name != "test" {
		t.Error("custom field label is not set properly")
	}
}

func TestGetBodyFromJson(t *testing.T) {
	t.Run("dataMap invalid type", func(t *testing.T) {
		var v Body
		err := GetBodyFromJson(Body{
			"test": make(chan int),
		}, &v)
		if err == nil {
			t.Error("should return error")
		}
	})

	t.Run("result is nil", func(t *testing.T) {
		var test *int = nil
		err := GetBodyFromJson(Body{
			"test": "ok",
		}, test)
		if err == nil {
			t.Error("should return error")
		}
	})

	t.Run("success", func(t *testing.T) {
		var v Body
		err := GetBodyFromJson(Body{
			"test": "ok",
		}, &v)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestDumpErrors(t *testing.T) {
	errs := ValidationErrors{
		"name": []error{
			errors.New("error 1"),
			errors.New("error 2"),
		},
	}
	dumpedErrs := DumpError(errs)
	if len(dumpedErrs) <= 0 {
		t.Error("result can not be empty")
	}
	nameErrs, ok := dumpedErrs["name"]
	if !ok {
		t.Error("result is not the same as input errors")
	}
	if len(nameErrs) != 2 {
		t.Error("result is not the same as input errors")
	}
	resultIsValid := true
	for _, nameErr := range nameErrs {
		if nameErr != "error 1" && nameErr != "error 2" {
			resultIsValid = false
			break
		}
	}
	if !resultIsValid {
		t.Error("result is not the same as input errors")
	}
}
