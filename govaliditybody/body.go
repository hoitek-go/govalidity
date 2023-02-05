package govaliditybody

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
)

var (
	IsTestMarshal         = false
	IsTestUnMarshal       = false
	IsTestParsedData      = false
	IsTestResultUnMarshal = false
)

func Bind(r *http.Request, result any) error {
	if r.PostForm == nil {
		r.ParseMultipartForm(32 << 20)
		r.ParseForm()
	}
	parsedData := map[string]interface{}{}
	bodyRaw, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if len(bodyRaw) > 0 {
		err = json.Unmarshal([]byte(bodyRaw), &parsedData)
		if err != nil {
			return err
		}
	} else {
		data, err := json.Marshal(r.Form)
		if IsTestMarshal {
			data, err = json.Marshal(math.Inf(1))
		}
		if err != nil {
			return err
		}
		dataMap := make(map[string][]interface{})
		err = json.Unmarshal(data, &dataMap)
		if IsTestUnMarshal {
			var test *int = nil
			err = json.Unmarshal(data, test)
		}
		if err != nil {
			return err
		}
		for key, value := range dataMap {
			parsedData[key] = value[0]
		}
	}
	byteArray, err := json.Marshal(parsedData)
	if IsTestParsedData {
		byteArray, err = json.Marshal(math.Inf(1))
	}
	if err != nil {
		return err
	}
	if IsTestResultUnMarshal {
		result = nil
	}
	err = json.Unmarshal(byteArray, result)
	if err != nil {
		return err
	}
	return nil
}
