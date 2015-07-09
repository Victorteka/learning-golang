package main

import "fmt"

// import "unicode/utf8"

func main() {
	letter := "A"
	alpha := "absdef"
	fmt.Println([]rune(letter)[0])
	fmt.Println([]rune(alpha)[2])
	fmt.Printf("%c\n", []rune(alpha)[2])

	// letter2 := "B"
	// buf := make([]byte, 3)
	// charCode := utf8.EncodeRune(buf, letter2)
	// fmt.Println(buf)
	// fmt.Println(charCode)

	for i, r := range "absdef" {
		if i == 2 {
			fmt.Println(r)
			break
		}
	}

	fmt.Println(charCodeAt("absdef", 2))
}

func charCodeAt(s string, n int) rune {
	for i, r := range s {
		if i == n {
			return r
		}
	}
	return 0
}
