package main

import ("fmt"; "time")

func main() {
  c := make(chan string)
  d := make(chan string)

  go func() {
    for{
      c <- "from 1"
      time.Sleep(time.Second * 2)
    }
  }()

  go func() {
    for{
      d <- "from 2"
      time.Sleep(time.Second * 3)
    }
  }()

  go func() {
    for{
      select{
        case msg1 := <- c:
          fmt.Println(msg1)
        case msg2 := <- d:
          fmt.Println(msg2)
        case <- time.After(time.Second):
          fmt.Println("timeout")
        default:
          fmt.Println("No channel is ready.")
      }
    }
  }()

  var input string
  fmt.Scanln(&input)
}
