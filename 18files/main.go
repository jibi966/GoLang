package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// only text files ( io, defer )
// other wise library

func main() {
	fmt.Println("Files in GoLang")
	content := "this needs to go in files"

	file, err := os.Create("./myFileGo.txt") // creation os operation

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is", length)

	defer file.Close()

	readFile("./myFileGo.txt")

}

func readFile(fileName string) { // reading ioutil
	byteData, err := ioutil.ReadFile(fileName)

	checkNilErr(err)

	fmt.Println("Text data inside file is\n", string(byteData))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
