package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"


	fmt.Println("List of all languages:" , languages)
	fmt.Println("JS:" , languages["JS"])
	fmt.Println("RB:" , languages["RB"])


	delete(languages, "RB")
	fmt.Println("List of all languages after removing ruby:" , languages)

	// loops interesting

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v", key, value)
	}

}