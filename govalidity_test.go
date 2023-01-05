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
	v := New("label")
	if v == nil {
		t.Error("new instance of validator struct can not be nil")
	}
}

func TestEmail(t *testing.T) {
	v := New("label")
	v.Email()
	if len(v.Validations) <= 0 {
		t.Error("Email validator should be set in validations")
	}
}

func TestRequired(t *testing.T) {
	v := New("label")
	v.Required()
	if len(v.Validations) <= 0 {
		t.Error("Required validator should be set in validations")
	}
}

func TestNumber(t *testing.T) {
	v := New("label")
	v.Number()
	if len(v.Validations) <= 0 {
		t.Error("Number validator should be set in validations")
	}
}

func TestUrl(t *testing.T) {
	v := New("label")
	v.Url()
	if len(v.Validations) <= 0 {
		t.Error("Url validator should be set in validations")
	}
}

func TestAlpha(t *testing.T) {
	v := New("label")
	v.Alpha()
	if len(v.Validations) <= 0 {
		t.Error("Alpha validator should be set in validations")
	}
}

func TestLowerCase(t *testing.T) {
	v := New("label")
	v.LowerCase()
	if len(v.Validations) <= 0 {
		t.Error("LowerCase validator should be set in validations")
	}
}

func TestUpperCase(t *testing.T) {
	v := New("label")
	v.UpperCase()
	if len(v.Validations) <= 0 {
		t.Error("UpperCase validator should be set in validations")
	}
}

func TestInt(t *testing.T) {
	v := New("label")
	v.Int()
	if len(v.Validations) <= 0 {
		t.Error("Int validator should be set in validations")
	}
}

func TestFloat(t *testing.T) {
	v := New("label")
	v.Float()
	if len(v.Validations) <= 0 {
		t.Error("Float validator should be set in validations")
	}
}

func TestJson(t *testing.T) {
	v := New("label")
	v.Json()
	if len(v.Validations) <= 0 {
		t.Error("Json validator should be set in validations")
	}
}

func TestIp(t *testing.T) {
	v := New("label")
	v.Ip()
	if len(v.Validations) <= 0 {
		t.Error("Ip validator should be set in validations")
	}
}

func TestIpV4(t *testing.T) {
	v := New("label")
	v.IpV4()
	if len(v.Validations) <= 0 {
		t.Error("IpV4 validator should be set in validations")
	}
}

func TestIpV6(t *testing.T) {
	v := New("label")
	v.IpV6()
	if len(v.Validations) <= 0 {
		t.Error("IpV6 validator should be set in validations")
	}
}

func TestPort(t *testing.T) {
	v := New("label")
	v.Port()
	if len(v.Validations) <= 0 {
		t.Error("Port validator should be set in validations")
	}
}

func TestIsDNSName(t *testing.T) {
	v := New("label")
	v.IsDNSName()
	if len(v.Validations) <= 0 {
		t.Error("IsDNSName validator should be set in validations")
	}
}

func TestHost(t *testing.T) {
	v := New("label")
	v.Host()
	if len(v.Validations) <= 0 {
		t.Error("Host validator should be set in validations")
	}
}

func TestLatitude(t *testing.T) {
	v := New("label")
	v.Latitude()
	if len(v.Validations) <= 0 {
		t.Error("Latitude validator should be set in validations")
	}
}

func TestLogitude(t *testing.T) {
	v := New("label")
	v.Logitude()
	if len(v.Validations) <= 0 {
		t.Error("Logitude validator should be set in validations")
	}
}

func TestAlphaNum(t *testing.T) {
	v := New("label")
	v.AlphaNum()
	if len(v.Validations) <= 0 {
		t.Error("AlphaNum validator should be set in validations")
	}
}

func TestInRange(t *testing.T) {
	v := New("label")
	v.InRange(0, 10)
	if len(v.Validations) <= 0 {
		t.Error("InRange validator should be set in validations")
	}
}

func TestMinMaxLength(t *testing.T) {
	v := New("label")
	v.MinMaxLength(0, 10)
	if len(v.Validations) <= 0 {
		t.Error("MinMaxLength validator should be set in validations")
	}
}

func TestMinLength(t *testing.T) {
	v := New("label")
	v.MinLength(0)
	if len(v.Validations) <= 0 {
		t.Error("MinLength validator should be set in validations")
	}
}

func TestMaxLength(t *testing.T) {
	v := New("label")
	v.MaxLength(10)
	if len(v.Validations) <= 0 {
		t.Error("MaxLength validator should be set in validations")
	}
}

func TestIn(t *testing.T) {
	v := New("label")
	v.In([]string{"test"})
	if len(v.Validations) <= 0 {
		t.Error("In validator should be set in validations")
	}
}

func TestCustomValidator(t *testing.T) {
	v := New("label")
	v.CustomValidator(func(s string, i ...interface{}) (bool, error) {
		return false, errors.New("test")
	})
	if len(v.Validations) <= 0 {
		t.Error("CustomValidator should be set in validations")
	}
}

func TestDefault(t *testing.T) {
	v := New("label")
	v.Default("test")
	if v.DefaultValue != "test" {
		t.Error("Default value is not set correctly")
	}
}

func TestValidateBody(t *testing.T) {
	t.Run("BodyIsNil", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		schema := Schema{
			"email":        New("email").Email().Required(),
			"name":         New("name").LowerCase().In([]string{"saeed", "taher"}).Required(),
			"age":          New("age").Number().Required(),
			"url":          New("url").Url().Required(),
			"json":         New("json").Json(),
			"ip":           New("ip").Ip().Required(),
			"filter[page]": New("filterpage").Int().InRange(10, 20).Required(),
		}
		type Query struct {
			Email string `json:"email"`
		}
		q := Query{}
		isValid, _ := ValidateBody(r, schema, &q)
		if isValid {
			t.Error("Should throw error when body is nil")
		}
	})

	t.Run("BodyIsNotPrepare", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", errReader(0))
		schema := Schema{
			"email":        New("email").Email().Required(),
			"name":         New("name").LowerCase().In([]string{"saeed", "taher"}).Required(),
			"age":          New("age").Number().Default("20").Required(),
			"url":          New("url").Url().Required(),
			"json":         New("json").Json(),
			"ip":           New("ip").Ip().Required(),
			"filter[page]": New("filterpage").Int().InRange(10, 20).Required(),
		}
		type Query struct {
			Email string `json:"email"`
		}
		q := Query{}
		isValid, _ := ValidateBody(r, schema, &q)
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
			"email":    New("email").Email().MinLength(5).Required(),
			"name":     New("name").LowerCase().MinMaxLength(3, 20).In([]string{"saeed", "taher"}).Required(),
			"lastName": New("lastname").LowerCase().MaxLength(25).Required(),
			"age":      New("age").Number().Default("20").Required(),
			"url":      New("url").Url().Required(),
			"json":     New("json").Json(),
			"ip":       New("ip").Ip().Required(),
			"filter[page]": New("filterpage").Int().InRange(10, 20).CustomValidator(func(s string, i ...interface{}) (bool, error) {
				return true, nil
			}).Required(),
		}
		type Query struct {
			Email string `json:"email"`
		}
		q := Query{}
		isValid, _ := ValidateBody(r, schema, &q)
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
			"email":    New("email").Email().MinLength(5).Required(),
			"name":     New("name").LowerCase().MinMaxLength(3, 20).In([]string{"saeed", "taher"}).Required(),
			"lastName": New("lastname").LowerCase().MaxLength(25).Required(),
			"age":      New("age").Number().Default("20").Required(),
			"url":      New("url").Url().Required(),
			"json":     New("json").Json(),
			"ip":       New("ip").Ip().Required(),
			"filter[page]": New("filterpage").Int().InRange(10, 20).CustomValidator(func(s string, i ...interface{}) (bool, error) {
				return false, errors.New("test")
			}).Required(),
		}
		type Query struct {
			Email string `json:"email"`
		}
		q := Query{}
		isValid, _ := ValidateBody(r, schema, &q)
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
			"email":    New("email").Email().MinLength(5).Required(),
			"name":     New("name").LowerCase().MinMaxLength(3, 20).In([]string{"saeed", "taher"}).Required(),
			"lastName": New("lastname").LowerCase().MaxLength(25).Required(),
			"age":      New("age").Number().Default("20").Required(),
			"url":      New("url").Url().Required(),
			"json":     New("json").Json(),
			"ip":       New("ip").Ip().Required(),
			"filter[page]": New("filterpage").Int().InRange(10, 20).CustomValidator(func(s string, i ...interface{}) (bool, error) {
				return false, errors.New("test")
			}).Required(),
		}
		type Query struct {
			Email string `json:"email"`
		}
		q := Query{}
		isValid, _ := ValidateBody(r, schema, &q)
		if isValid {
			t.Error("should be invalid")
		}
	})
}

func TestValidateQueries(t *testing.T) {
	t.Run("QueriesIsNotPrepare", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", errReader(0))
		schema := Schema{
			"email": New("email").Email().Required(),
		}
		type Query struct {
			Email string `json:"email"`
		}
		q := Query{}
		isValid, _ := ValidateQueries(r, schema, &q)
		if isValid {
			t.Error("Should throw error when body is not prepare")
		}
	})

	t.Run("QueriesIsPrepare", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/?email=value@email.com", errReader(0))
		schema := Schema{
			"email": New("email").Email().MinLength(5).Required(),
		}
		type Query struct {
			Email string `json:"email"`
		}
		q := Query{}
		isValid, _ := ValidateQueries(r, schema, &q)
		if !isValid {
			t.Error("should be valid")
		}
	})

	t.Run("tsetsetset", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, `/?email=sgh370@yahoo.com&filter={"phone":{"op":"equal","value":"09034005707"},"email":{"op":"equal","value":"sgh370@yahoo.com"},"name":{"op":"equal","value":"saeed"},"lastName":{"op":"equal","value":"ghanbari"},"userName":{"op":"equal","value":"sgh370"},"nationalCode":{"op":"equal","value":"0720464201"},"birthdate":{"op":"equal","value":"a"},"avatarUrl":{"op":"equal","value":"a"},"suspended_at":{"op":"equal","value":"a"},"created_at":{"op":"equal","value":"a"}}`, errReader(0))

		type FilterValue[T string | int] struct {
			Op    string `json:"op,omitempty"`
			Value T      `json:"value,omitempty"`
		}

		type UserFilterType struct {
			ID           FilterValue[int]    `json:"id,string,omitempty"`
			Phone        FilterValue[string] `json:"phone,omitempty"`
			Email        FilterValue[string] `json:"email,omitempty"`
			Name         FilterValue[string] `json:"name,omitempty"`
			LastName     FilterValue[string] `json:"lastName,omitempty"`
			UserName     FilterValue[string] `json:"userName,omitempty"`
			NationalCode FilterValue[string] `json:"nationalCode,omitempty"`
			Birthdate    FilterValue[string] `json:"birthdate,omitempty"`
			AvatarUrl    FilterValue[string] `json:"avatarUrl,omitempty"`
			SuspendedAt  FilterValue[string] `json:"suspended_at,omitempty"`
			CreatedAt    FilterValue[string] `json:"created_at,omitempty"`
		}

		type Query struct {
			Email  string         `json:"email,omitempty"`
			Filter UserFilterType `json:"filter,omitempty"`
		}

		q := &Query{}

		schema := Schema{
			"email": New("email").Email().Required(),
			"filter": Schema{
				"phone": Schema{
					"op":    New("filter.phone.op").Required(),
					"value": New("filter.phone.value").Required(),
				},
			},
		}

		isValid, _ := ValidateQueries(r, schema, q)
		if !isValid {
			t.Error("should be valid")
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

func TestDumpErrors(t *testing.T) {
	errs := ValidationErrors{
		"name": []error{
			errors.New("error 1"),
			errors.New("error 2"),
		},
	}
	dumpedErrs := DumpErrors(errs)
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
