package main

import "fmt"

// By using ... before the type name of the last parameter
// you can indicate that it takes zero or more of those parameters.
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	// fmt.Println(add(1,2,3,4,5,6,7))

	add(1, 2)
	add(1, 2, 3)

	// We can also pass a slice of ints by following the slice with ...
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(add(nums...))
}
