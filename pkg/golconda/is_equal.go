package golconda

import (
	"fmt"
	"reflect"
	"strings"
)

func IsEqual(field string, value interface{}) Operator {
	operator := Operator{}
	if value != nil {

		t := reflect.TypeOf(value).Kind()

		if t == reflect.Slice || t == reflect.Array {
			s := reflect.ValueOf(value)
			_vals := make([]interface{}, 0)
			valuesIn := buildQuestionMark(s.Len())
			operator.Expression = fmt.Sprintf("%s IN (%s)", field, valuesIn)
			for i := 0; i < s.Len(); i++ {
				_vals = append(_vals, s.Index(i))
			}
			operator.Vals = _vals
		} else {
			operator.Expression = fmt.Sprintf("%s = ?", field)
			_vals := make([]interface{}, 0)
			_vals = append(_vals, value)
			operator.Vals = _vals
		}

	}
	return operator
}

func buildQuestionMark(n int) string {
	s := make([]string, n)
	for i := range s {
		s[i] = "?"
	}
	fmt.Println(len(s), s)
	return strings.Join(s, ",")

}
