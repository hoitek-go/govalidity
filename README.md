# Go Validity
![Build Status](https://travis-ci.org/nock/nock.svg)
![Coverage Status](http://img.shields.io/badge/coverage-100%25-brightgreen.svg)
![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)

With this package you can validate your request easy peasy!

## Features 🔥
- Use Validation Easily In Your Model Based On Struct
- Custom Validation Support
- Nested Validation Support
- Router Queries Validation Support
- Router Params Validation Support
- i18n Support For Default Error Messages
- i18n Support For Field Names
- i18n Globally Support
- i18n Per Model Support
- Automatically Parse Body To Struct When Data is Valid
- Simple Usage
- 100% Test Coverage

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
	Email      string      `json:"email,omitempty"`
	Name       string      `json:"name,omitempty"`
	Age        json.Number `json:"age,omitempty"`
	Url        string      `json:"url,omitempty"`
	Json       string      `json:"json,omitempty"`
	Ip         string      `json:"ip,omitempty"`
	FilterPage json.Number `json:"filter[page],omitempty"`
}

func (data *UserIndexRequest) ValidateBody(r *http.Request) (bool, govalidity.ValidationErrors) {
	schema := govalidity.Schema{
		"email":        govalidity.New("email").Email().Required(),
		"name":         govalidity.New("name").LowerCase().In([]string{"saeed", "taher"}).Required(),
		"age":          govalidity.New("age").Number().Required(),
		"url":          govalidity.New("url").Url().Required(),
		"json":         govalidity.New("json").Json(),
		"ip":           govalidity.New("ip").Ip().Required(),
		"filter[page]": govalidity.New("filterPage").Int().InRange(10, 20).Required(),
	}
	return govalidity.ValidateBody(r, schema, &data)
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
		isValid, err := data.ValidateBody(r)
		if !isValid {
			log.Println(err)
			return
		}
		log.Printf("%#v\n", data)
	})
	http.ListenAndServe("127.0.0.1:9090", nil)
}
~~~ 

## Omit Fields When They Are Optional

If your schema has some fields which are optional, you can add omitempty value to json tag in the field.

For example:

~~~go  
type UserRequest struct {
	Email      string      `json:"email,omitempty"`
}
~~~ 

## Number Fields

If your schema has some fields which are number, you should add string value to json tag in the field.

For example:

~~~go  
type UserRequest struct {
	Age      int      `json:"email,string"`
}
~~~ 

## Nested Validation

You can use nested struct as validation. for example:

**Sample Schema:**

~~~go  
type FilterValue[T string | int] struct {
    Op    string `json:"op"`
    Value T      `json:"value"`
}

type FilterType struct {
    Name FilterValue[string] `json:"name"`
    Age  FilterValue[int]    `json:"age,omitempty"`
}

type Query struct {
    Email  string     `json:"email"`
    Filter FilterType `json:"filter"`
}

schema := govalidity.Schema{
    "email": New("email").Email().Required(),
    "filter": govalidity.Schema{
        "name": govalidity.Schema{
            "op":    New("filter.name.op").Alpha().FilterOperators().Required(),
            "value": New("filter.name.value").Alpha().Required(),
        },
        "age": govalidity.Schema{
            "op":    New("filter.age.op").Alpha().Required(),
            "value": New("filter.age.value").Alpha().Required(),
        },
    },
}
~~~ 

## Router Queries Validation

You can use govalidity for router queries. 

After you create the schema, just call ValidateQueries()

For example:

**Sample Schema:**

~~~go  
type Query struct {
    Email  string     `json:"email"`
}

func (data *Query) ValidateQueries(r *http.Request) (bool, govalidity.ValidationErrors) {
	schema := govalidity.Schema{
		"email":        govalidity.New("email").Email().Required(),
		"name":         govalidity.New("name").LowerCase().In([]string{"saeed", "taher"}).Required(),
		"age":          govalidity.New("age").Number().Required(),
		"url":          govalidity.New("url").Url().Required(),
		"json":         govalidity.New("json").Json(),
		"ip":           govalidity.New("ip").Ip().Required(),
	}
	return govalidity.ValidateQueries(r, schema, &data)
}
~~~ 

## Router Params Validation

You can use govalidity for router params. 

After you create the schema, just call ValidateParams()

For example:

**Sample Schema:**

~~~go  
type Query struct {
    Email  string     `json:"email"`
}

func (data *Query) ValidateParams(params govalidity.Params) (bool, govalidity.ValidationErrors) {
	schema := govalidity.Schema{
		"email":        govalidity.New("email").Email().Required(),
	}
	return govalidity.ValidateParams(params, schema, &data)
}
~~~ 

## Translate Default Error Messages

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
    "age": govalidity.New().Int().Default("20"),
}
~~~

## Optional Field

If you don't call Required() validator in schema field, your field can be optional.

So you can define optional fields in the schema like this:

~~~bash  
schema := govalidity.Schema{
    "age": govalidity.New().Int(),
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
AlphaNum() 
InRange(from, to int)
MinMaxLength(min, max int)
MinLength(min int)
MaxLength(max int)
In(in []string)
FilterOperators()
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
