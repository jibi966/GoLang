package main

import "fmt"

func main() {
	// this will be the structs or classes (we dont have classes in GoLang)

	// no inheritance in GoLang; No super or parent no child

	newUser := User{"Jibin", "jibin@go.dev", true, 23}
	fmt.Println(newUser)                      // {Jibin jibin@go.dev true 23}
	fmt.Printf("details are: %+v\n", newUser) // {Name:Jibin Email:jibin@go.dev Status:true Age:23} (need to give :=>  %+v)

	fmt.Printf("Name is %v and Email is %v", newUser.Name, newUser.Email) // Name is Jibin and Email is jibin@go.dev

}

// user export
type User struct {
	Name   string // fields can be access by any CAPS
	Email  string
	Status bool
	Age    int
}
