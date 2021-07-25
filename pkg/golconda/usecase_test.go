package golconda

import (
	"fmt"
	"testing"
)

type filters struct {
	byEmail    interface{}
	byName     interface{}
	byId       interface{}
	byLocation interface{}
}

// this is a real use case
func TestFullUseCase(t *testing.T) {
	// build the condition
	condition := NewAnd()

	// filters could come from func arguments or somewhere else
	filters := filters{}
	filters.byEmail = "myemail@test.it"
	filters.byId = []int{13, 21}
	filters.byLocation = 19

	condition.Append(
		IsEqual("email", filters.byEmail),
		// this trick is the scope of the
		IsEqual("name", filters.byName), // <nil>, so this will not be appended

		IsEqual("id", filters.byId), // Slice, so the equal operator is converted in IN operator
		IsEqual("location", filters.byLocation),
	)

	// "name" condition should not be printed
	expected := "(email = ? AND id IN (?,?) AND location = ?)"

	current := condition.Build()

	if expected != current {
		t.Errorf("Expected `%s`, got `%s`", expected, current)
	}

	vals := condition.Values()
	expectedLen := 4
	if len(vals) != expectedLen {
		t.Errorf("Expected length o condition.Values() to be %d, got %d", expectedLen, len(vals))
	}

	if vals[0] != filters.byEmail {
		t.Errorf("Expected vals[0] to be %s, got %s", filters.byEmail, vals[0])
	}

	if vals[1] != filters.byId.([]int)[0] {
		t.Errorf("Expected vals[1] to be %d, got %s", filters.byId.([]int)[0], vals[1])
	}

	if vals[2] != filters.byId.([]int)[1] {
		t.Errorf("Expected vals[2] to be %d, got %s", filters.byId.([]int)[1], vals[2])
	}

	if vals[3] != filters.byLocation {
		t.Errorf("Expected vals[3] to be %d, got %s", filters.byLocation, vals[3])
	}

	rawQuery := fmt.Sprintf("SELECT * FROM users WHERE %s", condition.Build())

	// ok this is pretty useless to test but...
	expectedRaw := "SELECT * FROM users WHERE " + expected

	if rawQuery != expectedRaw {
		t.Errorf("Raw query built failed. Expected `%s`, got `%s`  ", expectedRaw, rawQuery)
	}
	// and then...
	/*
		db.Raw(rawQuery, vals...)
	*/

}
