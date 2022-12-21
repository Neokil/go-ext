package slice

// Reduce takes a slice and uses the reduce-function to merge the first two elements, then
// continues to do so until only one element is left which is then returned.
// Example:
// ```
// arr := []int{1,2,3,4}
// sum := Reduce(arr, func(a int, b int) int { return a + b })
// ```
// This will sum up all elements in the array.
func Reduce[V any](a []V, reduce func(aValue V, bValue V) V) V {
	if len(a) == 0 {
		var result V

		return result
	}

	rValue := a[0]

	for i := 1; i < len(a); i++ {
		rValue = reduce(rValue, a[i])
	}

	return rValue
}
