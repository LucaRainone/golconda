package golconda

import (
	"fmt"
)

type Expression string
type operatorBuilder func(func() string) Operator

func SqlExpression(exp Expression) Expression {
	return exp
}

func genericOperator(field string, operatorString string, value interface{}) operatorBuilder {
	return func(paramPlaceholder func() string) Operator {
		operator := Operator{}

		if value != nil {
			switch value.(type) {
			case Expression:
				operator.Expression = fmt.Sprintf("%s %s %s", field, operatorString, value)
			default:
				operator.Expression = fmt.Sprintf("%s %s %s", field, operatorString, paramPlaceholder())
				operator.Vals = append(operator.Vals, value)
			}

		}

		return operator
	}
}

func genericExpression(field, expression string, flag bool) operatorBuilder {
	return func(paramPlaceholder func() string) Operator {
		operator := Operator{}
		if flag {
			operator.Expression = field + " " + expression
		}
		return operator
	}
}

func IsLess(field string, value interface{}) operatorBuilder {
	return genericOperator(field, "<", value)
}

func IsLessOrEqual(field string, value interface{}) operatorBuilder {
	return genericOperator(field, "<=", value)
}

func IsGreater(field string, value interface{}) operatorBuilder {
	return genericOperator(field, ">", value)
}

func IsGreaterOrEqual(field string, value interface{}) operatorBuilder {
	return genericOperator(field, ">=", value)
}

func IsLike(field string, value interface{}) operatorBuilder {
	return genericOperator(field, "LIKE", value)
}

func IsNotLike(field string, value interface{}) operatorBuilder {
	return genericOperator(field, "NOT LIKE", value)
}

func IsNull(field string, flag bool) operatorBuilder {
	return genericExpression(field, "IS NULL", flag)
}

func IsNotNull(field string, flag bool) operatorBuilder {
	return genericExpression(field, "IS NOT NULL", flag)
}
