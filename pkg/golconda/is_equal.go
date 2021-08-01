package golconda

import (
	"fmt"
	"reflect"
	"strings"
)

func isEqual(field string, value interface{}, isNot bool) operatorBuilder {
	return func(paramPlaceholder operatorParamBuilder) Operator {
		operator := Operator{}

		if value != nil {

			t := reflect.TypeOf(value).Kind()

			if t == reflect.Slice || t == reflect.Array {
				sqlOperator := "IN"
				if isNot {
					sqlOperator = "NOT IN"
				}

				s := reflect.ValueOf(value)
				if s.Len() == 0 {
					if isNot {
						operator.Expression = "TRUE"
					} else {
						operator.Expression = "FALSE"
					}
				} else {
					_vals := make([]interface{}, 0)

					valuesIn := buildQuestionMark(s.Len(), paramPlaceholder)
					operator.Expression = fmt.Sprintf("%s "+sqlOperator+" (%s)", field, valuesIn)

					for i := 0; i < s.Len(); i++ {
						_vals = append(_vals, s.Index(i).Interface())
					}

					operator.Vals = _vals
				}
			} else {
				sqlOperator := "="
				if isNot {
					sqlOperator = "!="
				}
				operator = genericOperator(field, sqlOperator, value)(paramPlaceholder)
			}

		}
		return operator
	}
}

func IsNotEqual(field string, value interface{}) operatorBuilder {
	return isEqual(field, value, true)
}

func IsEqual(field string, value interface{}) operatorBuilder {
	return isEqual(field, value, false)
}

func buildQuestionMark(n int, paramPlaceholder operatorParamBuilder) string {
	s := make([]string, n)
	for i := range s {
		s[i] = paramPlaceholder()
	}
	fmt.Println(len(s), s)
	return strings.Join(s, ",")

}
