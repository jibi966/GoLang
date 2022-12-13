package main

import "fmt"

func main() {

	jibin := User{"Jibin", "jibin@go.dev", true, 23}
	fmt.Println(jibin)
	jibin.GetStatus()
	jibin.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int // does not exportable if not capital
}

func (u User) GetStatus() {
	fmt.Println("Status is: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev" // does not modify original only copy, use pointer instead
	fmt.Println("Email of user is", u.Email)
}
