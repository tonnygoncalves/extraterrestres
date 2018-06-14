package equations

import "math"

// Calculate new x
func Calculate_X(d float64, v float64, i int) float64 {
	return d * math.Cos(((-v*math.Pi)/180)*float64(i))
}

// Calculate new y
func Calculate_Y(d float64, v float64, i int) float64 {
	return d * math.Sin(((-v*math.Pi)/180)*float64(i))
}
