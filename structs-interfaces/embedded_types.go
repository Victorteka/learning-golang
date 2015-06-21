package main

import "fmt"

type Person struct{
  name string
  age int
  height float64
}

type Android struct{
  Person
  model int
  id int
}

func (p *Person) greet() {
    fmt.Println("Hi, my name is", p.name)
}

func main() {
  me := Android{Person{"Android 3000", 25, 6.8}, 2015, 00101}
  me.greet()
}
