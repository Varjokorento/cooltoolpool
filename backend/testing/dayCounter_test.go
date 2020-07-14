package testing

import (
	"backend/dayCounter"
	"backend/structs"
	"testing"
)

func TestGetFieldsAsIntArray(t *testing.T) {
	var got = structs.NewDate(26, 2, 1993).GetFieldsAsIntArray()
	if got[0] != 26 && got[1] != 2 && got[2] != 1993 {
		t.Errorf("FieldArray = %d; want {26,2,1993}", got)
	}

}

func TestCreate(t *testing.T) {
	var dayOne = "26.02.1993"
	var dayTwo = "19.02.1993"
	var expected = structs.NewDate(26, 2, 1993)
	got := dayCounter.DayCount(dayOne, dayTwo)
	if got.GetFieldsAsIntArray() != expected.GetFieldsAsIntArray() {
		t.Errorf("DayCount(dayOne, dayTwo) = %d; want 26.02.1993", got)
	}
}
