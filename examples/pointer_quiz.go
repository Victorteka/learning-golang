package main
// 5. Write a program that can
// swap two integers (x := 1; y := 2; swap(&x, &y)
// should give you x=2 and y=1).

import "fmt"

func hotSwap(x *int, y *int) {
  // holds Y temperary so as to assign to X
  // Y is changed so X, so cant use Y again to assign
  // temp := *y
  // *y = *x
  // *x = temp
  // shortcut
  *x, *y = *y, *x
}

func main() {
  var (
    x = 3
    y = 5
  )
  hotSwap(&x, &y)
  fmt.Println(x, y)
}
