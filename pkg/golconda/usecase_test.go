package golconda

import (
	"fmt"
	"testing"
)

type filters struct {
	byEmail            interface{}
	byName             interface{}
	byId               interface{}
	byLocation         interface{}
	byDate             interface{}
	byDateStart        interface{}
	byDateEnd          interface{}
	byLastUpdateIsNull bool
	byConfirmDate      bool
	byLastUpdateStart  interface{}
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
	filters.byDate = SqlExpression("NOW()")
	filters.byDateStart = "2021-01-01"
	filters.byDateEnd = "2021-01-02"
	filters.byConfirmDate = true
	filters.byLastUpdateStart = "2021-02-02"

	condition.Append(
		IsEqual("email", filters.byEmail),
		IsEqual("name", filters.byName), // <nil>, so this will not be appended
		IsEqual("id", filters.byId),     // Slice, so the equal operator is converted in IN operator
		IsEqual("location", filters.byLocation),
		IsEqual("date", filters.byDate),
		IsBetween("date", filters.byDateStart, filters.byDateEnd),
		IsNull("last_update", filters.byLastUpdateIsNull),
		IsNotNull("confirm_date", filters.byConfirmDate),
		IsGreaterOrEqual("last_update", filters.byLastUpdateStart),
	)

	// "name" condition should not be printed
	expected := "(email = ? AND id IN (?,?) AND location = ? AND date = NOW() AND date BETWEEN ? AND ? AND confirm_date IS NOT NULL AND last_update >= ?)"

	current, vals := condition.Build()

	if expected != current {
		t.Errorf("Expected `%s`, got `%s`", expected, current)
	}

	expectedLen := 7
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

	if vals[4] != filters.byDateStart {
		t.Errorf("Expected vals[4] to be %d, got %s", filters.byDateStart, vals[4])
	}

	if vals[5] != filters.byDateEnd {
		t.Errorf("Expected vals[5] to be %d, got %s", filters.byDateEnd, vals[5])
	}

	if vals[6] != filters.byLastUpdateStart {
		t.Errorf("Expected vals[6] to be %d, got %s", filters.byLastUpdateStart, vals[6])
	}

	rawQuery := fmt.Sprintf("SELECT * FROM users WHERE %s", current)

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
