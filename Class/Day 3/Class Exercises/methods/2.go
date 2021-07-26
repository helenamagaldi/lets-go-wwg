package main

import (
	"fmt"
	"math"
)

// const pi = 3.14159265358979323846264338327950288419716939937510582097494459

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	circle1 := Circle{
		radius: 5,
	}
	area := circle1.Area()
	//fmt.Println(area)
	fmt.Printf("area: %.4f\n", area)
	perimeter := circle1.Perimeter()
	fmt.Printf("perimeter: %.4f\n", perimeter)
	fmt.Printf("circle:\n\tarea: %.4f\n\t, perimeter: %.4f\n\t, radius %.4f\n\t", area, perimeter, circle1.radius)
}
