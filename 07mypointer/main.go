package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of pointers")

	// creating a pointer
	// var ptr *int
	// fmt.Println("Value of the pointer is ", ptr)

	// Referencing a pointer
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of the actual pointer is ", ptr)
	fmt.Println("Value of the actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is ", myNumber)
	
}