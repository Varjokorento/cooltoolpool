package dayCounter

import (
	"backend/structs"
	"strconv"
	"strings"
)

// DayCount count the Days
func DayCount(dayOne string, dayTwo string) *structs.DateStar {
	var firstArray []string
	firstArray = splitDateIntoParts(dayOne)
	var dayAsString = firstArray[0]
	var monthAsString = firstArray[1]
	var yearAsString = firstArray[2]
	// TODO make the conversion
	/*var secondArray []string
	secondArray = splitDateIntoParts(dayTwo)*/

	var day int
	var month int
	var year int
	day, _ = strconv.Atoi(dayAsString)
	month, _ = strconv.Atoi(monthAsString)
	year, _ = strconv.Atoi(yearAsString)
	return structs.NewDate(day, month, year)
}

func splitDateIntoParts(day string) []string {
	var array []string = strings.Split(day, ".")
	return array
}
