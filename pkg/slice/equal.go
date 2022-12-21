package slice

// Equal takes two slices of comparables and returns true if all elements match.
func Equal[V comparable](a []V, b []V) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
