package main

import "fmt"

func main() {
	fmt.Println("If-Else GoLang")

	loginCount := 10

	var result string

	if loginCount < 10 {
		result = "Regular count"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Minimal"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}

	// web

	var checkNum int = 3

	if num := checkNum; num < 10 { // initialiZing value
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil {

	// }

}
