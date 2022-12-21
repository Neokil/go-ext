package float

import "math"

// Compare takes two floats and compares if they are equal up to the given precision
func Compare(f1 float64, f2 float64, precision int) bool {
	diff := math.Abs(f1 - f2)

	return diff < 1/math.Pow(10, float64(precision))
}
