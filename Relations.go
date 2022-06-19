package f64

import "math"

// Equivalent to: (math.IsNaN(fl1) && math.IsNaN(fl2)) || fl1 == fl2
func ValEquivalence(fl1, fl2 float64) bool {
	return fl1 == fl2 || (math.IsNaN(fl1) && math.IsNaN(fl2))
}
