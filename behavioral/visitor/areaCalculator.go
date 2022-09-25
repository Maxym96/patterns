package main

import (
	"fmt"
)

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(_ *Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(_ *Circle) {
	fmt.Println("Calculating area for circle")
}
func (a *AreaCalculator) visitForRectangle(_ *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
