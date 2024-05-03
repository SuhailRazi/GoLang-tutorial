package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()

	fmt.Println("unformatted Date ",presentTime)

	fmt.Println("formatted Date ",presentTime.Format("01-02-2006 03:04:05 Monday")) // for 24 hour clock format use 15:04:05

	createdDate := time.Date(2021, time.August, 20, 21, 20, 41, 0, time.UTC)

	fmt.Println("creteadDate ",createdDate)

	fmt.Println(" formatted created date", createdDate.Format("01-02-2006 15:04:05 Monday"))

}