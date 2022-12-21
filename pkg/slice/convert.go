package slice

import "fmt"

// Convert takes an array of Type K and converts it to an array of Type V using the convert-function on every element.
func Convert[K any, V any](a []K, convert func(K) V) []V {
	result := []V{}
	for _, ae := range a {
		result = append(result, convert(ae))
	}

	return result
}

// ConvertWithErr allows the user to specfy an error returning conversion function
// no partial conversion is allowed in this method.
func ConvertWithErr[K any, V any](a []K, convert func(K) (V, error)) ([]V, error) {
	result := []V{}
	for _, ae := range a {
		val, err := convert(ae)
		if err != nil {
			return nil, fmt.Errorf("converting slice failed: %w", err)
		}
		result = append(result, val)
	}

	return result, nil
}
