package slice

// Intersect returns all elements that exist in both arrays
// Example:
// ```
// arr1 := []int{1,2,3}
// arr2 := []int{2,3,4}
// intersect := Intersect(arr1, arr2)
// ```
// This will return an array containing {2,3}.
func Intersect[K comparable](a []K, b []K) []K {
	result := []K{}

	for _, ae := range a {
		for _, be := range b {
			if ae == be {
				result = append(result, ae)
			}
		}
	}

	return result
}

// IntersectCustom does the same as Intersect but with a custom compare function.
func IntersectCustom[K any, V any, P any](a []K, b []V, compare func(K, V) *P) []P {
	result := []P{}

	for _, ae := range a {
		for _, be := range b {
			r := compare(ae, be)
			if r != nil {
				result = append(result, *r)
			}
		}
	}

	return result
}
