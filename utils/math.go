package utils

// Max returns the maximum between two positive integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// BoundInt returns a value if it is in a interval, else the interval limit (lower or upper).
func BoundInt(n, minor, major int) int {
	if n < minor {
		return minor
	} else if n > major {
		return major
	} else {
		return n
	}
}

// BoundFloat returns a value if it is in a interval, else the interval limit (lower or upper).
func BoundFloat(n, minor, major float64) float64 {
	if n < minor {
		return minor
	} else if n > major {
		return major
	} else {
		return n
	}
}
