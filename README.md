# Go: The Good Parts

Go is an opinionated programming language initially developed by Google in 2007. It is a statically typed language with a simple syntax, resembling C or JavaScript. It features garbage collection, type safety and large standard library.

Go can be statically compiled into a single executable binary, which could target a large number of operating systems (from Linux and Windows to Plan 9) and processors (i386, amd64 and ARM).

```go
package main

// import "fmt"
// import "math"
// or
import (
  "math"
  "fmt"
)

// structs are used to define new data types
// shorter method
// type Circle struct{ x, y, r float64}
type Circle struct{
  x float64
  y float64
  r float64
}

// a method associated with a struct
func getCircleArea(ci *Circle) (float64) {
  return (math.Pi * ci.r * ci.r)
}

// equivalent to `void main` in C
func main() {
  // var c Cicle;
  // c := new(Circle) == {x: 0, y: 0, r: 0}
  // Set the default value if values are not defined
  // c := Circle{x: 0, y: 0, r: 5}
  // or just:
  c := Circle{0, 0, 5}

  area := getCircleArea(&c)
  fmt.Println(area)
  fmt.Println(c)
}
```

### Why a new language?

- C/C++ did not evolve with the computing landscape, no major systems language has emerged in over a decade: so there is a definite need for a new systems language, appropriate for needs of our computing era.

- In contrast to computing power, software development is not considerably faster or more successful (considering the number of failed projects) and applications still grow in size, so a new low-level language, but equipped with higher concepts, is needed.

- Before Go a developer had to choose between fast execution but slow and not efficient building (like C++), efficient compilation (but not so fast execution, like .NET or Java), or ease of programming (but slower execution, like the dynamic languages): Go is an attempt to combine all three wishes: efficient and thus fast compilation, fast execution, ease of programming.

### Uses of Go

- Go was originally conceived as a systems programming language, to be used in the heavy server-centric (Google) world of web servers, storage architecture and the like. For certain domains like high performance distributed systems Go has already proved to be a more productive language than most others. Go shines in and makes massive concurrency easy, so it should be a good fit for game server development.

- Complex event processing (CEP, see http://en.wikipedia.org/wiki/Complex_event_processing), where one needs both mass concurrency, high level abstractions and performance, is also an excellent target for Go usage. As we move to an Internet of Things, CEP will come to the forefront.

- But it turned out that it is also a general programming language, useful for solving text-processing problems, making frontends, or even scripting-like applications.
However Go is not suited for real-time software because of the garbage collection and automatic memory allocation.

- Go is used internally in Google for more and more heavy duty distributed applications; e.g. a part of Google Maps runs on Go.
Real life examples of usage of Go in other organizations can be found at http://go-lang.cat-v.org/ organizations-using-go. Not all uses of Go are mentioned there, because many companies consider this as private information. An application has been build inside Go for a large storage area network (SAN).(See Chapter 21 for a discussion of a sample of current use cases).

- A Go compiler exists for Native Client (NaCl) in the Chrome-browser; it will probably be used for the execution of native code in web applications in the Chrome OS.

- Go also runs on Intel as well as ARM processors (see chapter 2), so it runs under the Android OS, for example on Nexus smartphones.
