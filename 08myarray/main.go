package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list ", fruitList )
	fmt.Println("Fruit list length ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg list ", vegList )
	fmt.Println("Veg list length ", len(vegList))

}