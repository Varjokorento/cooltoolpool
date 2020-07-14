package dayCounter

import (
	"backend/structs"
	"strconv"
	"strings"
)

// DayCount count the Days
func DayCount(dayOne string, dayTwo string) *structs.DateStar {
	var firstArray []string
	firstArray = strings.Split(dayOne, ".")
	var dayAsString = firstArray[0]
	var monthAsString = firstArray[1]
	var yearAsString = firstArray[2]
	var day int
	var month int
	var year int
	day, _ = strconv.Atoi(dayAsString)
	month, _ = strconv.Atoi(monthAsString)
	year, _ = strconv.Atoi(yearAsString)
	return structs.NewDate(day, month, year)
}
