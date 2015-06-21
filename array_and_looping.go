package main

import "fmt"

func main() {
    // var x [5]float64 = { 98, 93, 77, 82, 83, }
    x := [5]float64{ 98, 93, 77, 82, 83 }
    getAverage(x);
}

func getAverage(x) {
     var total float64 = 0
     for i := 0; i < 5; i++ {
           total += x[i]
     }
     fmt.Println(total / 5)
}
