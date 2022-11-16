package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	length  int
	breadth int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {

	return r.length * r.breadth
}

func (c Circle) Area() float64 {

	return math.Pi * c.radius * c.radius
}

func main() {

	r := Rectangle{

		length:  10,
		breadth: 20,
	}

	c := Circle{

		radius: 10,
	}

	fmt.Printf("Area of Rectangle = %d\n", r.Area())

	fmt.Printf("Area of Circle = %f\n", c.Area())

}
