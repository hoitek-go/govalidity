package govaliditybody

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestBind(t *testing.T) {
	t.Run("WhenBodyIsPrepare", func(t *testing.T) {
		r := &http.Request{
			Form: url.Values{
				"name": []string{
					"test",
				},
			},
			Body: io.NopCloser(strings.NewReader("{\"name\":\"test\"}")),
		}
		var data *map[string]interface{} = &map[string]interface{}{}
		err := Bind(r, data)
		if err != nil {
			t.Error(err)
		}
		if data == nil {
			t.Error("body data can not be nil")
		}
		value, ok := (*data)["name"]
		if !ok {
			t.Error("body data is invalid")
		}
		if value != "test" {
			t.Error("body data value is invalid")
		}
	})

	t.Run("WhenBodyIsNotPrepare", func(t *testing.T) {
		r := &http.Request{
			Form: url.Values{
				"name": []string{
					"test",
				},
			},
			Body: io.NopCloser(strings.NewReader("")),
		}
		var data *map[string]interface{} = &map[string]interface{}{}
		err := Bind(r, data)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("WhenBodyIsInvalid", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", errReader(0))
		var data *map[string]interface{} = &map[string]interface{}{}
		err := Bind(r, data)
		if err == nil {
			t.Error("should return error when body has error")
		}
	})

	t.Run("WhenBodyUnmarshalHasError", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", io.NopCloser(strings.NewReader("s")))
		var data *map[string]interface{} = &map[string]interface{}{}
		err := Bind(r, data)
		if err == nil {
			t.Error("should return error when body is not json")
		}
	})

	t.Run("WhenFormIsReadyAndMarshalError", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		IsTestMarshal = true
		var data *map[string]interface{} = &map[string]interface{}{}
		err := Bind(r, data)
		IsTestMarshal = false
		if err == nil {
			t.Error("should return error")
		}
	})

	t.Run("WhenFormIsReadyAndUnmarshalError", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		IsTestUnMarshal = true
		var data *map[string]interface{} = &map[string]interface{}{}
		err := Bind(r, data)
		IsTestUnMarshal = false
		if err == nil {
			t.Error("should return error")
		}
	})

	t.Run("WhenFormIsReadyAndParsedDataError", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		IsTestParsedData = true
		var data *map[string]interface{} = &map[string]interface{}{}
		err := Bind(r, data)
		IsTestParsedData = false
		if err == nil {
			t.Error("should return error")
		}
	})

	t.Run("WhenFormIsReadyAndResultUnmarshalError", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		IsTestResultUnMarshal = true
		var data *map[string]interface{} = &map[string]interface{}{}
		err := Bind(r, data)
		IsTestResultUnMarshal = false
		if err == nil {
			t.Error("should return error")
		}
	})
}
