package golconda

import (
	"fmt"
	"testing"
)

func TestIsEqual(t *testing.T) {
	expected := "id = ?"
	operator := IsEqual("id", 42)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 1 {
		t.Errorf("Expected len of singleValue 1, got %d", len(operator.Vals))
	}

	if operator.Vals[0] != 42 {
		t.Errorf("Expected Values()[0] 42, got %s", operator.Vals[0])
	}

}

func TestIsEqualWithoutValue(t *testing.T) {
	expected := ""
	operator := IsEqual("id", nil)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len of singleValue 0, got %d", len(operator.Vals))
	}

}

func TestIsEqualWithArray(t *testing.T) {
	expected := "id IN (?,?)"
	operator := IsEqual("id", []int{13, 21})

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 2 {
		t.Errorf("Expected len of singleValue 2, got %d", len(operator.Vals))
	}

	if fmt.Sprint(operator.Vals[0]) != "13" {
		t.Errorf("Expected Values()[0] 13, got %d", operator.Vals[0])
	}

	if fmt.Sprint(operator.Vals[1]) != "21" {
		t.Errorf("Expected Values()[0] 21, got %d", operator.Vals[1])
	}
}
