package main

import (
	"math"
)

// Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle
type Circle struct {
	Radius float64
}

// Triangle
type Triangle struct {
	Base   float64
	Height float64
}

//Rectangle Methods

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// cale Methods (Pointer Receivers)
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

//Circle Methods

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// cale Methods (Pointer Receivers)
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

//Triangle Methods

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	// Per requirements: assuming an equilateral triangle
	return 3 * t.Base
}

// cale Methods (Pointer Receivers)
func (t *Triangle) Scale(factor float64) {
	t.Base *= factor
	t.Height *= factor
}
