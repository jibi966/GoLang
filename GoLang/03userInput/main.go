package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "this is simple welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza:")

	// comma ok syntax OR Err ok syntax
	input, err := reader.ReadString('\n')
	fmt.Println("this is input", input)
	fmt.Println("Err", err) // this will be <nil>

	fmt.Printf("type of input %T", input)

}
