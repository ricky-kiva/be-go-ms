package main

import (
	"fmt"
	"math"
)

type shapeFormula interface {
	area() float64
	circumference() float64
}

type circle struct {
	diameter float64
}

type square struct {
	side float64
}

// `(c circle)` means this function associated to `circle` struct
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius(), 2)
}

func (c circle) circumference() float64 {
	return math.Pi * c.diameter
}

func (sq square) area() float64 {
	return math.Pow(sq.side, 2)
}

func (sq square) circumference() float64 {
	return sq.side * 4
}

func (c circle) radius() float64 {
	return c.diameter / 2
}

func calculateShapes() {
	var formulas shapeFormula

	formulas = circle{3.1915}
	fmt.Printf("Circle -> Area: %.2f, Circumference: %2.f\n", formulas.area(), formulas.circumference())

	formulas = square{10}
	fmt.Printf("Square -> Area: %.2f, Circumference: %.2f\n", formulas.area(), formulas.circumference())
}
