package govalidityo

import "encoding/json"

type OperatorValue struct {
	Op    string
	Value interface{}
}

const (
	EQUALS              = "equals"
	NOT_EQUALS          = "neq"
	GREATER_THAN        = "gt"
	GREATER_THAN_EQUALS = "gte"
	LESS_THAN           = "lt"
	LESS_THAN_EQUAL     = "lte"
	LIKE                = "like"
	NOT_LIKE            = "nlike"
	IGNORE_LIKE         = "ilike"
	NOT_IGNORE_LIKE     = "nilike"
	IN                  = "in"
	NOT_IN              = "nin"
	IS                  = "is"
	IS_NOT              = "isnot"
	BETWEEN             = "between"
	NOT_BETWEEN         = "nbetween"
	OVERLAP             = "overlap"
	CONTAINS            = "contains"
	CONTAINED           = "contained"
	STARTS_WITH         = "startsWith"
	ENDS_WITH           = "endsWith"
	IS_EMPTY            = "isEmpty"
	IS_NOT_EMPTY        = "isNotEmpty"
	IS_ANY_OF           = "isAnyOf"
)

var MapSqlOperators = map[string]string{
	EQUALS:              "=",
	NOT_EQUALS:          "!=",
	GREATER_THAN:        ">",
	GREATER_THAN_EQUALS: ">=",
	LESS_THAN:           "<",
	LESS_THAN_EQUAL:     "<=",
	LIKE:                "LIKE",
	NOT_LIKE:            "NOT LIKE",
	IGNORE_LIKE:         "iLIKE",
	NOT_IGNORE_LIKE:     "NOT iLIKE",
	IN:                  "IN",
	NOT_IN:              "NOT IN",
	IS:                  "=",
	IS_NOT:              "<>",
	BETWEEN:             "IN",
	NOT_BETWEEN:         "NOT IN",
	OVERLAP:             "=",
	CONTAINS:            "LIKE",
	CONTAINED:           "LIKE",
	STARTS_WITH:         "LIKE",
	ENDS_WITH:           "LIKE",
	IS_EMPTY:            "IS NULL",
	IS_NOT_EMPTY:        "IS NOT NULL",
	IS_ANY_OF:           "IN",
}

func getJsonSlice(str string) []string {
	var strSlice []string
	isValid := json.Unmarshal([]byte(str), &strSlice) == nil
	if isValid {
		return strSlice
	}
	return []string{}
}

func GetSqlOperatorValue(govalidityOperator string, value string) *OperatorValue {
	op, ok := MapSqlOperators[govalidityOperator]
	if !ok {
		op = "="
	}
	var val interface{} = value
	switch govalidityOperator {
	case LIKE, NOT_LIKE, IGNORE_LIKE, NOT_IGNORE_LIKE:
		val = "%" + value + "%"
	case IN, NOT_IN:
		val = getJsonSlice(value)
	}
	return &OperatorValue{
		Op:    op,
		Value: val,
	}
}
