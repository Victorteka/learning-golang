package main

import "fmt"

const name string = "Eugene Mutai"

// this is a comment
func main(int i = 3) {
    var a string = "Hello world!!"
    b := "Do not stay at home."

    fmt.Println(a + b)
    fmt.Println(true && false)

    switch i {
        case 0: fmt.Println("Zero")
        case 1: fmt.Println("One")
        case 2: fmt.Println("Two")
        case 3: fmt.Println("Three")
        case 4: fmt.Println("Four")
        case 5: fmt.Println("Five")
        default: fmt.Println("Unknown Number")
    }
}

