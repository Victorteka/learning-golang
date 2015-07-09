package main

import (
	console "fmt"
	"math"
)

type Circle struct {
	radius, area float64
}

type Rectangle struct {
	length, width, area float64
}

type Square struct {
	Rectangle
}

type Shapes interface {
	Area() float64
}

type AllShapes []Shapes

func (c Circle) Area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (a AllShapes) area() {
	var area float64
	for _, s := range a {
		area += s.Area()
	}

	console.Println("Total Areas of Allshapes", area)
}

func totalAreas(shapes ...Shapes) {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}

	console.Println("Total Areas", area)
}

func main() {
	c := Circle{radius: 7.0}
	// c.Area()
	// console.Println("This is the area of the circle: ", c.area)

	r := Rectangle{10, 13, 0}
	// r.Area()
	// console.Println("This is the area of the rectangle: ", r.area)

	totalAreas(r, c)

	shapes := make(AllShapes, 2)
	shapes[0], shapes[1] = r, c
	shapes.area()
}
