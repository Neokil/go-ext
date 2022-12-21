package slice

// Any returns true if any of the elements of the array passes the check.
func Any[K any](a []K, check func(K) bool) bool {
	for _, ae := range a {
		if check(ae) {
			return true
		}
	}

	return false
}
