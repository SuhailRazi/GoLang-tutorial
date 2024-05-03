package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello")
	myDefer()
}

// world one two

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
