package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println("welcom to pizza app")

	// fmt.Println("Please rate our pizza betweeen 1 to 5")

	// reader := bufio.NewReader(os.Stdin)

	// input, _ := reader.ReadString('\n')

	// fmt.Println("Thanks for rating", input)

	// // strconv used to convert the variables
	// numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Added 1 to rating", numRating+1)
	// }

	// mode := bufio.NewReader(os.Stdin)

	// name, _ := mode.ReadString('\n')

	// fmt.Println("this is name", name)

	// // convert to string
	// stringedValue, newErr := strconv.ParseFloat(strings.TrimSpace(name), 64)

	// if newErr != nil {
	// 	fmt.Println("this  is Err", newErr)
	// } else {
	// 	fmt.Println("this is the value", stringedValue)
	// }

	const giveNum string = "948"

	numberValue, ErrNumber := strconv.ParseFloat(strings.TrimSpace(giveNum), 64)

	if ErrNumber != nil {
		fmt.Println("Errror while converting", ErrNumber)
	} else {
		fmt.Println("String converted", numberValue+1)
	}

	fmt.Println(giveNum)
}
