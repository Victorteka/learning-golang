package main

import (
  "math"
  "fmt"
)

type Circle struct{ x, y, r float64}
type Square struct{ x, y float64}

// DO not name the functions since the method will be shared
// across all shapes
func (ci *Circle) area() float64 {
  return (math.Pi * ci.r * ci.r)
}

// square function
func (sq *Square) area() float64 {
  return (sq.x * sq.x)
}

// ------------ INTERFACES --------------
// Used to implement relationships shared by two or more structures (structs)
// In our case both Square and Circle have area methods which return float64s
// so both types implement the Shape interface.
type Shape interface{
  area() float64
}

func getShapeArea(shapes ...Shape) float64 {
  var area float64
  for _, shape := range shapes {
    area += shape.area()
  }
  return area
}

func main() {
  c := Circle{0, 0, 5}
  s := Square{4, 4}

  fmt.Println(c.area(), s.area())
  fmt.Println(getShapeArea(&c, &s))
}
