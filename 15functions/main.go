package main

import "fmt"

func main() { // this is the entry point
	fmt.Println("Function In GoLang")
	greeter()

	// func greeterTwo(){ // not allowed to write functions inside function
	// 	fmt.Println("Hai from 2")
	// }

	result := adder(3, 5)
	fmt.Println("Result is", result)

	res, msg := proAdded(2, 5, 8, 7)
	fmt.Println("new result is", res)
	fmt.Println("new msg is", msg)

}

func greeter() {
	fmt.Println("Namasthe JS")
}

func adder(a int, b int) int {
	return a + b
}

// func (){ // unknown functions exists
// 	fmt.Println("possible")
// }

func proAdded(values ...int) (int, string) {
	var total int = 0

	for _, value := range values {
		total += value
	}

	// for i := range values {
	// 	total += values[i]
	// }

	return total, "Hi result function"
}
