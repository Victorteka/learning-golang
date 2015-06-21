package main

import (
  "math"
  "fmt"
)

type Circle struct{ x, y, r float64}

// ------------------ METHODS -------------
// In between the keyword func and the name of the function we've added a “receiver
// by creating the function in this way it allows us to call the function using the . operator
func getCircleArea(ci *Circle) area() (float64) {
  return (math.Pi * ci.r * ci.r)
}

func main() {
  c := Circle{0, 0, 5}

  // Now C inherits the method area
  // Class like inheritance in JS but automated
  // Like magic!!!
  area := c.area()
  fmt.Println(area)
  fmt.Println(c)
}
