package structs

type DateStar struct {
	day   int `json:"day"`
	month int `json:"month"`
	year  int `json:"year"`
}

func NewDate(day int, month int, year int) *DateStar {
	d := DateStar{day: day, month: month, year: year}
	return &d
}
