package slice

// Subtract returns the elements in a that are not in b.
func Subtract[K comparable](a []K, b []K) []K {
	difference := []K{}
	for _, elementA := range a {
		if !Any(b, func(elementB K) bool { return elementB == elementA }) {
			difference = append(difference, elementA)
		}
	}

	return difference
}
