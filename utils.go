package main

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func checkIntMinimumValue(n, min int) bool {
	return n > min
}

func checkFloatMinimumValue(n, min float64) bool {
	return n > min
}

func createFillArray(length, value int) []int {
	array := make([]int, length)
	for i := 0; i < length; i++ {
		array[i] = value
	}
	return array
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func boundInt(n, minor, major int) int {
	if n < minor {
		return minor
	} else if n > major {
		return major
	} else {
		return n
	}
}
