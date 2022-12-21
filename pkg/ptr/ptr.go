package ptr

// Create will take any value and return a new pointer to that value.
func Create[K any](v K) *K {
	return &v
}

// Resolve will take a Pointer and return the Value or the default value of the underlying type.
func Resolve[K any](v *K) K {
	if v == nil {
		var result K

		return result
	}

	return *v
}

// ResolveWithFallback will take a Pointer and return the Value or the provided fallback value.
func ResolveWithFallback[K any](v *K, fallback K) K {
	if v == nil {
		return fallback
	}

	return *v
}
