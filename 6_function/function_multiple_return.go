package main

import "math"

func calculateCircle(d float64) (float64, float64) {
	var area = math.Pi * math.Pow((d/2), 2)
	var circumference = math.Pi * d

	return area, circumference
}
