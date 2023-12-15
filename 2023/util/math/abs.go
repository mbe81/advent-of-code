package math

// Abs returns te positive value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
