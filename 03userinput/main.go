package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user User"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for out pizza:")

	// Comma ok syntax || error ok 

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)


}