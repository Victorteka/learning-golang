package main
// --------- POINTERS ---------
// Pointers reference a location in memory where a value is stored
// rather than the value itself. (They point to something else)
// By using a pointer (*int) the zero function is able to modify the original variable.

import "fmt"

func makeZero(point_back_to_var *int) {
     *point_back_to_var = 0
}

func main() {
     x := 5
     makeZero(&x)
     fmt.Println(x) // x is 0
}
