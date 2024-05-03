package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	greeter()

	result := adder(3, 5)

	fmt.Println("Result is", result)

	proResult, myMessage := proAdded(2, 5, 8, 7, 4)
	fmt.Println("ProResult is", proResult, myMessage)

}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdded(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "String result"
}

func greeter() {
	fmt.Println("Namastey from golang")
}
