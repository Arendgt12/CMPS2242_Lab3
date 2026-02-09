package main

import (
	"math"
	"testing"
)

//Area Tests

func TestRectangleArea(t *testing.T) {
	r := Rectangle{Width: 10, Height: 5}
	got := r.Area()
	want := 50.0
	if got != want {
		t.Errorf("Rectangle.Area() = %.2f; want %.2f", got, want)
	}
}

func TestCircleArea(t *testing.T) {
	c := Circle{Radius: 10}
	got := c.Area()
	want := math.Pi * 100
	if got != want {
		t.Errorf("Circle.Area() = %.2f; want %.2f", got, want)
	}
}

func TestTriangleArea(t *testing.T) {
	tr := Triangle{Base: 10, Height: 5}
	got := tr.Area()
	want := 25.0
	if got != want {
		t.Errorf("Triangle.Area() = %.2f; want %.2f", got, want)
	}
}

//Perimeter Tests

func TestRectanglePerimeter(t *testing.T) {
	r := Rectangle{Width: 10, Height: 5}
	got := r.Perimeter()
	want := 30.0
	if got != want {
		t.Errorf("Rectangle.Perimeter() = %.2f; want %.2f", got, want)
	}
}

func TestCirclePerimeter(t *testing.T) {
	c := Circle{Radius: 10}
	got := c.Perimeter()
	want := 2 * math.Pi * 10
	if got != want {
		t.Errorf("Circle.Perimeter() = %.2f; want %.2f", got, want)
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tr := Triangle{Base: 10}
	got := tr.Perimeter()
	want := 30.0 // Equilateral: 3 * 10
	if got != want {
		t.Errorf("Triangle.Perimeter() = %.2f; want %.2f", got, want)
	}
}

// Scale Tests

func TestRectangleScale(t *testing.T) {
	r := Rectangle{Width: 10, Height: 5}
	r.Scale(2)
	if r.Width != 20 || r.Height != 10 {
		t.Errorf("Rectangle.Scale(2) failed. Got Width=%.2f, Height=%.2f", r.Width, r.Height)
	}
}

func TestCircleScale(t *testing.T) {
	c := Circle{Radius: 10}
	c.Scale(0.5)
	if c.Radius != 5 {
		t.Errorf("Circle.Scale(0.5) failed. Got Radius=%.2f", c.Radius)
	}
}

func TestTriangleScale(t *testing.T) {
	tr := Triangle{Base: 10, Height: 4}
	tr.Scale(3)
	if tr.Base != 30 || tr.Height != 12 {
		t.Errorf("Triangle.Scale(3) failed. Got Base=%.2f, Height=%.2f", tr.Base, tr.Height)
	}
}
