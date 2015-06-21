package main

import ("fmt"; "time")

// A channel that doesn't have these restrictions is known as bi-directional.
// A bi-directional channel can be passed to a function
// that takes send-only or receive-only channels,
// but the reverse is not true.
func pinger(c chan string) {
  for i := 0; ;i++ {
    c <- "ping"
  }
}

func ponger(c chan string) {
  for i:=0; ;i++ {
    c <- "pong"
  }
}

// <-chan signifies that printer can only receive messages
func printer(c <-chan string) {
  for{
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan string = make(chan string)
  go pinger(c)
  go ponger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}
