package main

import "fmt"

// can not declare no-var types here
var jwtToken = 800

// public first letter CAPS
const LoginToken string = "this is the string"

func main() {
	var userName string = "jibin"
	fmt.Println(userName)
	fmt.Printf("Variable is typeof: %T \n", userName)

	var isLoaggedIn bool = false
	fmt.Println(isLoaggedIn)
	fmt.Printf("Variable is typeof: %T \n", isLoaggedIn)

	var smallVal int = 256

	// many integer types are there
	fmt.Println(smallVal)
	fmt.Printf("Variable is typeof: %T \n", smallVal)

	var smallValF float64 = 256.5555555

	// many integer types are there
	fmt.Println(smallValF)
	fmt.Printf("Variable is typeof: %T \n", smallValF)

	// default values and aliases
	var anotherVar int // 0
	fmt.Println(anotherVar)
	fmt.Printf("Variable is typeof: %T \n", anotherVar)

	// implicit type
	var website = "google.com"
	const we = "aa"
	// website = 0
	fmt.Println(website, we)

	// no var style only inside method
	numberOfUser := 800
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Println(jwtToken)

}
