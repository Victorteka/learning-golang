package main
// The Fibonacci sequence is defined as:
// fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2)...

import "fmt"

func fib(num int) int {
  if num == 0 {
    return 0
  } else if num == 1 {
    return 1
  }

  return fib(num-1) + fib(num-2)
}

func main() {
  fmt.Print("Enter a number: ")

  var input float64
  fmt.Scanf("%f", &input)

  // x := 2
  fmt.Println(fib(int(input)))
}
