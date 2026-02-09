package main

import (
	"fmt"
)

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}
	tri := Triangle{Base: 6, Height: 4}

	fmt.Println("--- Initial Shapes ---")
	fmt.Printf("Rectangle: Area = %.2f, Perimeter = %.2f\n", rect.Area(), rect.Perimeter())
	fmt.Printf("Circle:    Area = %.2f, Perimeter = %.2f\n", circ.Area(), circ.Perimeter())
	fmt.Printf("Triangle:  Area = %.2f, Perimeter = %.2f\n", tri.Area(), tri.Perimeter())

	fmt.Println("\n--- After Scaling ---")
	rect.Scale(2.0)
	fmt.Printf("Scaled Rectangle (2x): Width = %.2f, Height = %.2f\n", rect.Width, rect.Height)
	fmt.Printf("New Rectangle Area:    %.2f\n", rect.Area())
}
