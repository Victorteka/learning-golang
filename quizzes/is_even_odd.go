package main

import "fmt"

func main() {
  isEvenOrOdd();
}

func isEvenOrOdd() {
    for i := 1; i <= 10; i++ {
        if i % 2 == 0 {
            fmt.Println(i, "even")
        } else {
            fmt.Println(i, "odd")
        }
    }
}
