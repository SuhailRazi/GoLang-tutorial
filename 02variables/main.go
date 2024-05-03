package main

import "fmt"

const LoginToken string = "token" // Uppercase by convention, publich declaration

func main() {
	var username string = "Suhail"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type : %T \n", smallValue)

	var smallFloat float64 = 255.455445545554555
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	// Default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// implicit type for declaring variable

	var website = "Hellowebsite.in"
	fmt.Println(website)

	// No var style

	numberOfUser := 30000
	fmt.Println(numberOfUser)


	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}