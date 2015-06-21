package main

import ("fmt"; "time"; "math/rand")


func con(num int) {
  for i := 0; i < 10; i++ {
    fmt.Println(num, ":", i)
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
  }
}

func main() {
  for i := 0; i < 10; i++ {
    go con(i)
  }

  var input string
  fmt.Scanln(&input)
}
