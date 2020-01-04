package iteration

// Repeat character for "repeatCount" times the received parameter
func Repeat(character string, repeatCount int) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
