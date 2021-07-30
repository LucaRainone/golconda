package golconda

import (
	"fmt"
)

type Expression string

func SqlExpression(exp Expression) Expression {
	return exp
}

func genericOperator(field string, operatorString string, value interface{}) Operator {
	operator := Operator{}

	if value != nil {
		switch value.(type) {
		case Expression:
			operator.Expression = fmt.Sprintf("%s %s %s", field, operatorString, value)
		default:
			operator.Expression = fmt.Sprintf("%s %s ?", field, operatorString)
			operator.Vals = append(operator.Vals, value)
		}

	}

	return operator
}

func genericExpression(field, expression string, flag bool) Operator {
	operator := Operator{}
	if flag {
		operator.Expression = field + " " + expression
	}
	return operator
}

func IsLess(field string, value interface{}) Operator {
	return genericOperator(field, "<", value)
}

func IsLessOrEqual(field string, value interface{}) Operator {
	return genericOperator(field, "<=", value)
}

func IsGreater(field string, value interface{}) Operator {
	return genericOperator(field, ">", value)
}

func IsGreaterOrEqual(field string, value interface{}) Operator {
	return genericOperator(field, ">=", value)
}

func IsLike(field string, value interface{}) Operator {
	return genericOperator(field, "LIKE", value)
}

func IsNotLike(field string, value interface{}) Operator {
	return genericOperator(field, "NOT LIKE", value)
}

func IsNull(field string, flag bool) Operator {
	return genericExpression(field, "IS NULL", flag)
}

func IsNotNull(field string, flag bool) Operator {
	return genericExpression(field, "IS NOT NULL", flag)
}
