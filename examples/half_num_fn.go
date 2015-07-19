package main

// Write a function which takes an integer and halves it
// and returns true if it was even or false if it was odd.
// For example half(1) should return (0, false)
// and half(2) should return (1, true).

import "fmt"

func half(num int) (int, bool) {
	half := num / 2
	if num%2 == 0 {
		return half, true
	} else {
		return half, false
	}
}

func main() {
	fmt.Print("Enter a number: ")

	var input float64
	fmt.Scanf("%f", &input)

	// x := 2
	fmt.Println(half(int(input)))
}
