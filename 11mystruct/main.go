package main

import "fmt"

func main() {
	fmt.Println("Struct in Go")

	// No inheritance, no super, parent

	suhail := User{"Suhail", "Suhail@go.dev", true, 16}
	fmt.Println(suhail)

	fmt.Printf("Suhail details are : %+v \n", suhail)
	fmt.Printf("Name is %v and email is %v", suhail.Name, suhail.Email)
}

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age 	int
}