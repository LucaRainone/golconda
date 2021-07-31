package golconda

import (
	"testing"
)

func TestIsBetween(t *testing.T) {
	expected := "date BETWEEN ? AND ?"

	from := "2021-10-01"
	to := "2022-10-31"

	operator := IsBetween("date", from, to)(func() string { return "?" })

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 2 {
		t.Errorf("Expected len of operator.Vals 2, got %d", len(operator.Vals))
	}

	if operator.Vals[0] != from {
		t.Errorf("Expected len of operator.Vals[0] %s, got %s", from, operator.Vals[0])
	}
	if operator.Vals[1] != to {
		t.Errorf("Expected len of operator.Vals[1] %s, got %s", to, operator.Vals[1])
	}

}

func TestIsBetweenWithoutFrom(t *testing.T) {
	expected := "date <= ?"

	to := "2022-10-31"
	operator := IsBetween("date", nil, to)(func() string { return "?" })

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 1 {
		t.Errorf("Expected len of operator.Vals 1, got %d", len(operator.Vals))
	}

	if operator.Vals[0] != to {
		t.Errorf("Expected len of operator.Vals[0] %s, got %s", to, operator.Vals[0])
	}

}

func TestIsBetweenWithoutTo(t *testing.T) {
	expected := "date >= ?"

	from := "2021-10-01"
	operator := IsBetween("date", from, nil)(func() string { return "?" })

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 1 {
		t.Errorf("Expected len of operator.Vals 1, got %d", len(operator.Vals))
	}

	if operator.Vals[0] != from {
		t.Errorf("Expected len of operator.Vals[0] %s, got %s", from, operator.Vals[0])
	}

}
