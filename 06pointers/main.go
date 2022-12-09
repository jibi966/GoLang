package main

import "fmt"

func main() {
	fmt.Println("WELCOME Pointers")

	// creating pointer
	// var ptr *int
	// fmt.Println("Pointer", ptr)

	// creating pointer which is referncing to that value => &
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("new pointer is", ptr)  // this will give the memory address
	fmt.Println("new pointer is", *ptr) // this will give the actual value

	// operation will only performed in original value
	*ptr = *ptr + 2
	fmt.Println("new value is", myNumber)

}
