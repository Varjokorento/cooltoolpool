package gpaCounter

func GetGpa(credits int, sumOfGrades int) float32 {
	return float32(sumOfGrades) / float32(credits)
}
