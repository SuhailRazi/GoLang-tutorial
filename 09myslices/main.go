package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the video of slices")

	var fruitList = []string{"Apple", "Orange", "Mango"}
	// fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Peach", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	// fmt.Println("slice after append",fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 455
	highScore[3] = 566
	// highScore[4] = 777 brings error since exceeds limit

	highScore = append(highScore, 777, 888, 999)
	// fmt.Println(highScore)

	sort.Ints(highScore) //Sorting
	fmt.Println("sorted",highScore)
	// fmt.Println("check ",sort.IntsAreSorted(highScore)) //sorting check

	// How to remove a value from slices based on index

	var courses = []string{"Java", "Python", "Golang", "C", "C++"}
	fmt.Println(courses)
	var index int = 2
	courses= append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}