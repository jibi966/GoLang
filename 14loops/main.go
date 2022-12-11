package main

import "fmt"

func main() {

	// Loops in GoLang
	// days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	// type 1
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// type 2
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// type 3
	// for index, value := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, value)
	// }

	// type 4 like while loop
	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 { // goto this will break the loop and go to the declared condition
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
			// break
		}

		fmt.Println("value is ", rougueValue)
		rougueValue++ // there is no ++rougueValue thing in Go
	}

lco:
	fmt.Println("Jumbing at here while hitting 2")

}
