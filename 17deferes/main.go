package main

import "fmt"

func main() {
	defer fmt.Println("World") // execute at the end of the function.
	defer fmt.Println("One")
	// defer myDefer() // this is a defer
	defer fmt.Println("Two") // the last defer will execute first
	fmt.Println("Hello")
	myDefer() // this is not a defer
}

func myDefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i) // this is defer, last will execute first ( LIFO )
	}
}
