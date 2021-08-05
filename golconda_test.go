package golconda

import (
	"testing"
)

func TestAndEmptyCondition(t *testing.T) {
	expected := "(TRUE)"
	c := NewAnd()
	emptyCondition, _ := c.Build()
	if emptyCondition != expected {
		t.Errorf("Expected %s, got %s", expected, emptyCondition)
	}

}

func TestOrEmptyCondition(t *testing.T) {
	expected := "(FALSE)"
	c := NewOr()
	emptyCondition, _ := c.Build()
	if emptyCondition != expected {
		t.Errorf("Expected %s, got %s", expected, emptyCondition)
	}
}

func TestAppend(t *testing.T) {
	expected := "(somecondition)"
	c := NewAnd()
	operator := Operator{}
	operator.Expression = "somecondition"
	operator.Vals = append(operator.Vals, 1)

	c.Append(func(paramPlaceholder operatorParamBuilder) Operator { return operator })

	conditionString, values := c.Build()
	if conditionString != expected {
		t.Errorf("Expected %s, got %s", expected, conditionString)
	}

	if values[0] != 1 {
		t.Errorf("values: Expected %d, got %d", operator.Vals[0], values[0])
	}
}

func TestAppendCondition(t *testing.T) {
	expected := "(somecondition AND (anothercondition OR other))"
	c := NewAnd()
	operator := Operator{}
	operator.Expression = "somecondition"

	sub := NewOr()
	operatorSub := Operator{}
	operatorSub.Expression = "anothercondition OR other"
	operatorSub.Vals = append(operatorSub.Vals, 1)
	sub.Append(func(paramPlaceholder operatorParamBuilder) Operator { return operatorSub })

	c.Append(func(paramPlaceholder operatorParamBuilder) Operator { return operator }, sub.AsOperator())

	conditionString, _ := c.Build()
	if conditionString != expected {
		t.Errorf("Expected %s, got %s", expected, conditionString)
	}

}
