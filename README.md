# Go Validity
![Build Status](https://travis-ci.org/nock/nock.svg)
![Coverage Status](http://img.shields.io/badge/coverage-100%25-brightgreen.svg)
![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)

With this package you can validate your request easy peasy!

## Features 🔥
- Use validation easily in your model based on struct
- Custom validation support
- i18n support for default error messages
- i18n support for field names
- i18n globally support
- i18n per model support
- Automatically parse body to struct when data is valid
- Simple usage
- 100% test coverage

## Installation 
Run the following comand in your project

~~~bash  
  go get github.com/hoitek-go/govalidity
~~~

## Usage/Examples  
**Sample Model file:**

user-index-request.go
~~~go  
package models

import (
	"encoding/json"
	"net/http"

	"github.com/hoitek-go/govalidity"
	"github.com/hoitek-go/govalidity/govaliditym"
)

type UserIndexRequest struct {
	Email      string      `json:"email"`
	Name       string      `json:"name"`
	Age        json.Number `json:"age"`
	Url        string      `json:"url"`
	Json       string      `json:"json"`
	Ip         string      `json:"ip"`
	FilterPage json.Number `json:"filter[page]"`
}

func (b *UserIndexRequest) ValidateBody(r *http.Request) (bool, govalidity.ValidationErrors, govalidity.Body) {
	schema := govalidity.Schema{
		"email":        govalidity.New().Email().Required(),
		"name":         govalidity.New().LowerCase().In([]string{"saeed", "taher"}).Required(),
		"age":          govalidity.New().Number().Required(),
		"url":          govalidity.New().Url().Required(),
		"json":         govalidity.New().Json(),
		"ip":           govalidity.New().Ip().Required(),
		"filter[page]": govalidity.New().Int().InRange(10, 20).Required(),
	}
	return govalidity.ValidateBody(r, schema)
}
~~~ 

**Sample Main File:**

main.go
~~~go  
package main

import (
	"log"
	"net/http"

	"github.com/hoitek-go/govalidity"
	"myproject/models"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := &models.UserIndexRequest{}
		isValid, err, json := data.ValidateBody(r)
		if !isValid {
			log.Println(err)
			return
		}
		bodyErr := govalidity.GetBodyFromJson(json, data)
		if bodyErr != nil {
			log.Println(bodyErr)
			return
		}
		log.Printf("%#v\n", data)
	})
	http.ListenAndServe("127.0.0.1:9090", nil)
}
~~~ 

## Translate Default Error Messages:

You can change default error messages globally or in your model.

For example if you want to change the IsEmail validation error message, you should put the following codes in your codes:

~~~go  
govalidity.SetDefaultErrorMessages(&govaliditym.Validations{
    IsEmail: "{field} can be valid email",
})
~~~ 

**Common Attributes In Translation**

You can use the following attributes in all error messages:
~~~go  
{field} // show the name of the field in the error message 
{value} // show the value of the user's input in the error message
~~~  

You can use the following attributes in Range() validator:
~~~go  
{from} // show the minimum value of Range() validator in the error message 
{to} // show the maximum value of Range() validator in the error message
~~~  

You can use the following attributes in MinMaxLength() validator:
~~~go  
{min} // show the minimum value of MinMaxLength() validator in the error message 
{max} // show the maximum value of MinMaxLength() validator in the error message
~~~  

You can use the following attributes in MinLength() validator:
~~~go  
{min} // show the minimum value of MinLength() validator in the error message
~~~  

You can use the following attributes in MaxLength() validator:
~~~go  
{max} // show the maximum value of MaxLength() validator in the error message
~~~  

You can use the following attributes in In() validator:
~~~go  
{in} // show the list of acceptable values from In() validator in the error message
~~~  

## Translate Schema Fields

You can change the field labels at any area of your codes may be globally or in the model file per model.

For example we want to change the label of "name" and "filter[page]" keys to "first name" and "page". 

To do that you should use the following codes:

~~~bash  
govalidity.SetFieldLabels(&govaliditym.Labels{
    "name":         "first name",
    "filter[page]": "page",
})
~~~

## Custom Validator

You can define your own validator easily.

For example we want to define a custom validator to check if the name is "saeed".

To do that you can define you schema like this:

~~~bash  
func CheckName(field string, params ...interface{}) (bool, error) {
	if params[0].(string) != "saeed" {
		return false, errors.New("{field} can be \"saeed\", but you send {value}")
	}
	return true, nil
}
schema := govalidity.Schema{
    "name": govalidity.New().CustomValidator(CheckName).Required(),
}
~~~

## Default Value

You can set default value for each fields in the schema like this:
~~~bash  
schema := govalidity.Schema{
    "name": govalidity.New().Int().Default("20"),
}
~~~

## List of Validators

~~~bash  
Email() 
Required() 
Number() 
Url() 
Alpha() 
LowerCase() 
UpperCase() 
Int() 
Float() 
Json() 
Ip() 
IpV4() 
IpV6() 
Port() 
IsDNSName() 
Host() 
Latitude() 
Logitude() 
AlpaNum() 
InRange(from, to int)
MinMaxLength(min, max int)
MinLength(min int)
MaxLength(max int)
In(in []string)
~~~

## Run Tests

~~~bash  
  make test
~~~

## Export Test Coverage

~~~bash  
  make testcov
~~~

## Tech Stack  
**Server:** Golang
 
## Licence  
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)   
