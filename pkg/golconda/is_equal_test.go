package golconda

import (
	"fmt"
	"testing"
)

func TestIsEqual(t *testing.T) {
	expected := "id = ?"
	operator := IsEqual("id", 42)(func() string { return "?" })

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
	operator := IsEqual("id", nil)(func() string { return "?" })

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len of singleValue 0, got %d", len(operator.Vals))
	}

}

func TestIsEqualWithArray(t *testing.T) {
	expected := "id IN (?,?)"
	operator := IsEqual("id", []int{13, 21})(func() string { return "?" })

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

func TestIsEqualWithEmptyArray(t *testing.T) {
	expected := "FALSE"
	operator := IsEqual("id", []int{})(func() string { return "?" })
	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}
	if len(operator.Vals) != 0 {
		t.Errorf("Expected len of singleValue 0, got %d", len(operator.Vals))
	}
}

func TestIsNotEqual(t *testing.T) {
	expected := "id != ?"
	operator := IsNotEqual("id", 42)(func() string { return "?" })

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

func TestIsNotEqualWithoutValue(t *testing.T) {
	expected := ""
	operator := IsNotEqual("id", nil)(func() string { return "?" })

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len of singleValue 0, got %d", len(operator.Vals))
	}

}

func TestIsNotEqualWithArray(t *testing.T) {
	expected := "id NOT IN (?,?)"
	operator := IsNotEqual("id", []int{13, 21})(func() string { return "?" })

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

func TestIsNotEqualWithEmptyArray(t *testing.T) {
	expected := "TRUE"
	operator := IsNotEqual("id", []int{})(func() string { return "?" })
	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}
	if len(operator.Vals) != 0 {
		t.Errorf("Expected len of singleValue 0, got %d", len(operator.Vals))
	}
}
