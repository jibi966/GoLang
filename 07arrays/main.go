package main

import "fmt"

func main() {
	fmt.Println("WELCOME TO ARRAYS IN GOLang")

	var fruitList [4]string // basic declare

	// if we dont give fruitList[2] it will give an empty space in the output array
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Peach"
	fmt.Println("LIST", fruitList)
	fmt.Println("LIST Length", len(fruitList))

	var vegetablelist = [5]string{"potato", "beans", "shroom"} // even we filled only 3 values the length will be 5 always as we initialized
	fmt.Println("Veggie is", vegetablelist)
	fmt.Println("Veggie is", len(vegetablelist))

	// var arr = [3]int{1, 2, 3}
	// var arr [2]string
	// arr[0] = "a"

}
