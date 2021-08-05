package golconda

import (
	"testing"
)

func questionMarkBuilder() string {
	return "?"
}

func TestIsLess(t *testing.T) {
	expected := "field < ?"

	value := 123
	operator := IsLess("field", value)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 1 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 1, len(operator.Vals))
	}

	if operator.Vals[0] != value {
		t.Errorf("Expected %s, got %s", operator.Vals[0], operator.Expression)
	}
}

func TestIsLessOrEqual(t *testing.T) {
	expected := "field <= ?"

	value := 1234
	operator := IsLessOrEqual("field", value)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 1 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 1, len(operator.Vals))
	}

	if operator.Vals[0] != value {
		t.Errorf("Expected %s, got %s", operator.Vals[0], operator.Expression)
	}
}

func TestIsGreater(t *testing.T) {
	expected := "field > ?"

	value := 321
	operator := IsGreater("field", value)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 1 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 1, len(operator.Vals))
	}

	if operator.Vals[0] != value {
		t.Errorf("Expected %s, got %s", operator.Vals[0], operator.Expression)
	}
}

func TestIsGreaterOrEqual(t *testing.T) {
	expected := "field >= ?"

	value := 911
	operator := IsGreaterOrEqual("field", value)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 1 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 1, len(operator.Vals))
	}

	if operator.Vals[0] != value {
		t.Errorf("Expected %s, got %s", operator.Vals[0], operator.Expression)
	}
}

func TestIsLike(t *testing.T) {
	expected := "field LIKE ?"

	value := "abc"
	operator := IsLike("field", value)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 1 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 1, len(operator.Vals))
	}

	if operator.Vals[0] != value {
		t.Errorf("Expected %s, got %s", operator.Vals[0], operator.Expression)
	}
}

func TestIsNotLike(t *testing.T) {
	expected := "field NOT LIKE ?"

	value := "acb"
	operator := IsNotLike("field", value)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 1 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 1, len(operator.Vals))
	}

	if operator.Vals[0] != value {
		t.Errorf("Expected %s, got %s", operator.Vals[0], operator.Expression)
	}
}

func TestIsNull(t *testing.T) {
	expected := "field IS NULL"

	operator := IsNull("field", true)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 0, len(operator.Vals))
	}

}

func TestIsNotNull(t *testing.T) {
	expected := "field IS NOT NULL"

	operator := IsNotNull("field", true)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 0, len(operator.Vals))
	}
}

func TestIsLessIgnored(t *testing.T) {
	expected := ""

	operator := IsLess("field", nil)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 0, len(operator.Vals))
	}

}

func TestIsLessOrEqualIgnored(t *testing.T) {
	expected := ""

	operator := IsLessOrEqual("field", nil)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 0, len(operator.Vals))
	}

}

func TestIsGreaterIgnored(t *testing.T) {
	expected := ""

	operator := IsGreater("field", nil)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 0, len(operator.Vals))
	}

}

func TestIsGreaterOrEqualIgnored(t *testing.T) {
	expected := ""

	operator := IsGreaterOrEqual("field", nil)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 0, len(operator.Vals))
	}

}

func TestIsLikeIgnored(t *testing.T) {
	expected := ""

	operator := IsLike("field", nil)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 0, len(operator.Vals))
	}

}

func TestIsNotLikeIgnored(t *testing.T) {
	expected := ""

	operator := IsNotLike("field", nil)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 0, len(operator.Vals))
	}

}

func TestIsNullIgnored(t *testing.T) {
	expected := ""

	operator := IsNull("field", false)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 0, len(operator.Vals))
	}

}

func TestIsNotNullIgnored(t *testing.T) {
	expected := ""

	operator := IsNotNull("field", false)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 0, len(operator.Vals))
	}
}

func TestExpression(t *testing.T) {
	expected := "field < NOW()"

	var value Expression = "NOW()"
	operator := IsLess("field", value)(questionMarkBuilder)

	if operator.Expression != expected {
		t.Errorf("Expected %s, got %s", expected, operator.Expression)
	}

	if len(operator.Vals) != 0 {
		t.Errorf("Expected len(operator.Vals) to be %d, got %d", 0, len(operator.Vals))
	}

}
