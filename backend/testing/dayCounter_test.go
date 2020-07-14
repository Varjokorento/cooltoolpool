package testing

import (
	"backend/dayCounter"
	"backend/structs"
	"testing"
)

func TestParse(t *testing.T) {
	var dayOne = "26.02.1993"
	var dayTwo = "19.02.1993"
	var expected = structs.NewDate(1, 1, 2020)
	got := dayCounter.DayCount(dayOne, dayTwo)
	if got != expected {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
