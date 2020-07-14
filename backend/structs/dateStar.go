package structs

type DateStar struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

func NewDate(day int, month int, year int) *DateStar {
	d := DateStar{Day: day, Month: month, Year: year}
	return &d
}

func (r *DateStar) CreateDate(day int, month int, year int) *DateStar {
	d := DateStar{Day: day, Month: month, Year: year}
	return &d
}

func (d DateStar) GetFieldsAsIntArray() [3]int {
	return [3]int{d.Day, d.Month, d.Year}
}
