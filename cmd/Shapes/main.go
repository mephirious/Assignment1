package main

import . "github.com/NursultanNurgaliyev/Assignment1/Shapes/models"

func main() {
	shapes := []Shape{
		Rectangle{Length: 10, Width: 10},
		Circle{Radius: 1},
		Square{Length: 30},
		Triangle{SideA: 10, SideB: 10, SideC: 10},
	}

	for _, shape := range shapes {
		PrintDetailsShape(shape)
	}
}
