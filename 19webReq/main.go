package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const Url = "https://fakestoreapi.com/products"

func main() {
	fmt.Println("Handling web Req")

	response, err := http.Get(Url) // get method

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
	fmt.Printf("type is %T", response)

	defer response.Body.Close() // callers responsibility to close the connection

	dataBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println(content)

}
