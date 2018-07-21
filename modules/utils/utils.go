package utils

// Primes is a function that calculates primes up to limit value
func Primes(limit int) []int {
	// Key control
	var divisibles int
	results := make([]int, 1)
	for i := 1; i < limit; i++ {
		divisibles = 0
		for j := 1; j < limit; j++ {
			if (i % j) == 0 {
				divisibles++
			}
		}
		if divisibles == 2 {
			results = append(results, i)
		}
	}
	return results
}
