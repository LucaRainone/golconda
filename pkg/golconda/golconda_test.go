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

	c.Append(func(paramPlaceholder func() string) Operator { return operator })

	conditionString, values := c.Build()
	if conditionString != expected {
		t.Errorf("Expected %s, got %s", expected, conditionString)
	}

	if values[0] != 1 {
		t.Errorf("values: Expected %d, got %d", operator.Vals[0], values[0])
	}
}
