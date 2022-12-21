package slice

// Contains returns true if the array contains the element..
func Contains[K comparable](a []K, v K) bool {
	for _, ae := range a {
		if ae == v {
			return true
		}
	}

	return false
}
