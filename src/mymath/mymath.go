package mymath

import (
	"math"
)

func SolveQuadratic(a float64, b float64, c float64) (bool, float64, float64) {
	bbm4ac := (b * b) - (4 * a * c)
	if bbm4ac < 0 {
		return false, 0.0, 0.0
	} else {
		root := math.Sqrt(bbm4ac)
		sol1 := ((-b) + root) / (2 * a)
		sol2 := ((-b) - root) / (2 * a)
		return true, sol1, sol2
	}
}