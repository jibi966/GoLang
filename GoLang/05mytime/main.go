package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time GoLang")

	// time lib
	presentTime := time.Now()
	fmt.Println("the time is", presentTime)

	// this is the default Time,Date,Day format for Go
	fmt.Println("new formatted time", presentTime.Format("01-02-2006 15:04:05 Monday"))

	// creating date
	createdDate := time.Date(2020, time.December, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println("created time", createdDate)

	// if we want to format we need to give this default variables NO MATTER WHAT
	fmt.Println("created time formate", createdDate.Format("01-02-2006 15:04:05 Monday"))

	// present := time.Now()
	// present.Format("01-02-2006 Monday")

}
