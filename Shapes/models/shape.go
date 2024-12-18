package models

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

func PrintDetailsShape(s Shape) {
	fmt.Printf("Area: %10.2f, Perimeter: %10.2f\n", s.Area(), s.Perimeter())
}
