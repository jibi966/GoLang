package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Orange", "Main"} // if you use this syntax to create slice you need to initialize
	fmt.Printf("Type is %T\n", fruitList)

	// fruitList[0] => array

	var newfruitList = append(fruitList, "Tomato", "Mango") // use to add items to slices
	fmt.Println(fruitList)
	fmt.Println(newfruitList)

	// var slicedfruitList = append(newfruitList[2:3]) // used to slice the slices from index 2 to 3 (will not included)
	// var slicedfruitList = append(newfruitList[:3]) start from 0 to 3
	// var slicedfruitList = append(newfruitList[2:]) start from 2 to end
	// fmt.Println(slicedfruitList)

	highScores := make([]int, 4) // way to make slices

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 didn't work

	highScores = append(highScores, 555, 666, 321) // we can add this way

	fmt.Println(highScores)
	sort.Ints(highScores) // will sort in ascending order
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // will give boolean if the slice is sorted

}
