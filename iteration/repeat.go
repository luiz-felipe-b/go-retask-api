package iteration

// Repeat takes a character and repeats it 5 times
func Repeat(character string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += character
	}
	return
}
