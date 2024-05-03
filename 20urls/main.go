package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.youtube.com/learn?courseName=reactjs&paymentId=fshdbiw"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myUrl)

	// Parsing
	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("The types of query params is : %T \n", qParams)

	fmt.Println(qParams["courseName"])

	for _, val := range qParams {
		fmt.Println("Params is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "/tutCss",
		RawPath: "user=suhail",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)

}
