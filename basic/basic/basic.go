package basic

import "math"

func calcTriangle(a int, b int) int {

	c := 0
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}
