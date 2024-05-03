package main

import "fmt"

func main() {
	fmt.Println("Struct in Go")

	// No inheritance, no super, parent

	suhail := User{"Suhail", "Suhail@go.dev", true, 16}
	fmt.Println(suhail)

	fmt.Printf("Suhail details are : %+v \n", suhail)
	fmt.Printf("Name is %v and email is %v \n", suhail.Name, suhail.Email)
	suhail.GetStatus()
	suhail.AddEmail()

	fmt.Println("The opriginal email is ", suhail.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("The status of the user is \n", u.Status)
}

func (u User) AddEmail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is : ", u.Email)
}
