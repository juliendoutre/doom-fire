package grid

import "fmt"

type errMinimumValue struct {
	name       string
	lowerBound int
}

func (e *errMinimumValue) Error() string {
	return fmt.Sprintf("%s should be at least %d", e.name, e.lowerBound)
}
