package main

import (
	"fmt"
)

func main() {
	fmt.Println("Channels in go lang")

	myCh := make(chan int)

	go func(ch chan int) {}(myCh)

	// fmt.Println(<-myCh)
	// myCh <- 5

}
