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

func TestAndSingleCondition(t *testing.T) {
	expected := "(id = ?)"
	c := NewAnd()
	c.IsEqual("id", "42")
	singleCondition := c.Build()
	singleValue := c.Values()
	if singleCondition != expected {
		t.Errorf("Expected %s, got %s", expected, singleCondition)
	}

	if len(singleValue) != 1 {
		t.Errorf("Expected len of singleValue 1, got %d", len(singleValue))
	}

	if singleValue[0] != "42" {
		t.Errorf("Expected Values()[0] 42, got %s", singleValue[0])
	}

}
