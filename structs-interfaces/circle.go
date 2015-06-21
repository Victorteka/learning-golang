package main
// ------------------ STRUCTS -------------
// They are used to define new data types
// EG. Circles

// import "fmt"
// import "math"
import (
  "math"
  "fmt"
)

type Circle struct{
  x float64
  y float64
  r float64
}
// shorter method
// type Circle struct{ x, y, r float64}

func getCircleArea(ci *Circle) (float64) {
  return (math.Pi * ci.r * ci.r)
}

func main() {
  // var c Cicle;
  // c := new(Circle) == {x: 0, y: 0, r: 0}
  // Set the default value if values are not defined
  // c := Circle{x: 0, y: 0, r: 5}
  c := Circle{0, 0, 5}

  area := getCircleArea(&c)
  fmt.Println(area)
  fmt.Println(c)
}
