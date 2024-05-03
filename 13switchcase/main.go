package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch and case in golang")

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You an move 2 spots")
	case 3:
		fmt.Println("You an move 3 spots")
	case 4:
		fmt.Println("You an move 4 spots")
		fallthrough //to get the next case also
	case 5:
		fmt.Println("You an move 5 spots")
	case 6:
		fmt.Println("You an move 6 spots and roll dice again")
	default:
		fmt.Println("What was that")
	}
}
