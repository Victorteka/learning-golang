package main

import "fmt"

func main() {

	const goLang string = "GO Language"

	// ----------- NOTE: ---------------
	// Do not declare values that you will not use
	// The compiler fails

	// my name
	// var name string = "Eugene Mutai"
	name := "Eugene Mutai"

	// age
	var age int = 25

	// float
	var howHandsome float64 = 8.75

	// arrays
	fruits := [5]string{"apple", "oranges", "pinapple", "pears", "peaches"}
	count := [3]int{1, 2, 3}

	for index, fruit := range fruits {
		fmt.Println(fruit)
	}

	// object
	aboutme := make(map[string]string)
	aboutme["name"] = "Eugene Mutai"
	// aboutme["age"] = "25"
	// aboutme["location"] = "Woodley, Kilimani"

	// The other way of declaring an object
	// always has a comma in the last declaration
	// unlike JavaScript
	moreAboutMe := map[string]string{
		"name": "Eugene Mutai",
		//  "age": "25",
		//  "location": "Woodley, Kilimani",
	}

	// For multidimentional object checking
	// Age will not be printed if it does not exist
	if age, ok := aboutme["age"]; ok {
		fmt.Println(age)
	}

	fmt.Println("Learning" + count)
	fmt.Println(name, age, howHandsome, fruits)
	fmt.Println(aboutme["name"], moreAboutMe["name"])
	fmt.Println(count)

}
