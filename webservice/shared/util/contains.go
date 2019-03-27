package util

// Contains func
func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}

	return false
}
