package main

import "fmt"

func main() {
	// how to remove a value from slices based on index
	var courses = []string{"ReactJs", "JavaScript", "swift", "Python", "Ruby"}
	fmt.Println(courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
