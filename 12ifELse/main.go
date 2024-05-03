package main

import "fmt"

func main() {
	fmt.Println("If else in Golang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular use"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Result is 10"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("Num is Not less than 10")
	}
}