package main

import (
	"fmt"
	"regexp/syntax"
)

func main() {

	str := []rune("beta")   // use rune slice
	acharacter := rune('a') // use single quote, instead of double quote
	onerune := rune('吃')

	fmt.Printf("%v \n", string(str)) // convert back to string

	fmt.Printf("%v \n", string(acharacter)) // convert back to string

	fmt.Printf("%v \n", string(onerune)) // convert back to string

	// there are times when accessing str is not acceptable because
	// it is a slice. Therefore, you just have to reference the first
	// element

	// for example :

	ok := syntax.IsWordChar(str[0]) // won't work without [0]

	fmt.Printf("%v is a word ? : %v \n", string(str), ok)

}
