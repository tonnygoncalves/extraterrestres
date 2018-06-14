package equations

import (
	"math"
)

// CalculateX get new x
func CalculateX(d float64, v float64, i int) float64 {
	return d * math.Cos(((-v*math.Pi)/180)*float64(i))
}

// CalculateY get new y
func CalculateY(d float64, v float64, i int) float64 {
	return d * math.Sin(((-v*math.Pi)/180)*float64(i))
}

// RoundingInt the number
func RoundingInt(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// RoundingDecimal with a number of decimals
func RoundingDecimal(num float64, precition int) float64 {
	output := math.Pow(10, float64(precition))
	return float64(RoundingInt(num*output)) / output
}

// distanceBetween2points: calculate the distance betweet two points
func distanceBetween2points(X1 float64, X2 float64, Y1 float64, Y2 float64) float64 {
	var result = math.Pow((X2-X1), 2) + math.Pow((Y2-Y1), 2)
	result = math.Pow(result, 0.5)
	return RoundingDecimal(result, 8)
}
