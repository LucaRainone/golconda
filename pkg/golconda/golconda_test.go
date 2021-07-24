package golconda

import (
	"testing"
)

func TestAndEmptyCondition(t *testing.T) {
	expected := "(TRUE)"
	c := NewAnd()
	emptyCondition := c.Build()
	if emptyCondition != expected {
		t.Errorf("Expected %s, got %s", expected, emptyCondition)
	}

}

func TestOrEmptyCondition(t *testing.T) {
	expected := "(FALSE)"
	c := NewOr()
	emptyCondition := c.Build()
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

	c.Append(operator)

	conditionString := c.Build()
	if conditionString != expected {
		t.Errorf("Expected %s, got %s", expected, conditionString)
	}

	if operator.Vals[0] != 1 {
		t.Errorf("Vals: Expected %d, got %d", 1, operator.Vals[0])
	}
}
