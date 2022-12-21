package slice

// Where takes a slice and returns a new slice only containing the elements for which the check-function returns true.
func Where[K any](a []K, check func(K) bool) []K {
	result := []K{}

	for _, ae := range a {
		if check(ae) {
			result = append(result, ae)
		}
	}

	return result
}
