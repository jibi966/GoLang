package main

import "fmt"

func main() {
	// fmt.Println("Maps in GOLang")

	languages := make(map[string]string) // make to create map
	// ang := make(map[int]string)

	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println(languages)
	fmt.Println("JS:", languages["JS"])
	fmt.Println("RB:", languages["RB"])
	fmt.Println("PY:", languages["PY"])

	delete(languages, "RB") // to delete
	fmt.Println("After delete:", languages)

	// loops in maps GoLang

	for key, value := range languages { // key may not be neede as in Err,Ok syntax ( _ , value) := is used maily in such cases
		fmt.Printf("key is %v and value is %v\n", key, value)
	}

}
