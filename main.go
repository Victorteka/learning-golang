// This is know as a “package declaration”.
// Every Go pro- gram must start with a package declaration.
// Packages are Go's way of organizing and reusing code.
// There are two types of Go programs: executables and libraries.
package main

// The import keyword is how we include code from other packages to use with our program.
// The fmt package (shorthand for format) implements formatting for in- put and output.
// The use of double quotes like this is known as a “string literal” which is a type of “expression”.
import "fmt"

// this is a comment
// All functions start with the keyword func 
// followed by the name of the function (main in this case)
func main(int i = 3) {
    var a string = "Hello world!!"
    b := "Do not stay at home."

    fmt.Println(a + b)
    fmt.Println(true && false)
}
