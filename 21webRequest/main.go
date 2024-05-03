package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to the web verb video")
	// const myUrl string = "http://localhost:8000/get"
	const myUrl string = "http://localhost:8000/post"
	const myUrl2 string = "http://localhost:8000/postform"

	// PerformGetRequest(myUrl)
	PerformPostJsonRequest(myUrl)
	PerformFormRequest(myUrl2)
}

// Get functionality
// func PerformGetRequest(myUrl string) {

// 	response, err := http.Get(myUrl)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer response.Body.Close()

// 	fmt.Println("Status code: ", response.StatusCode)
// 	fmt.Println("Content length:", response.ContentLength)

// 	var responseString strings.Builder

// 	content, _ := io.ReadAll(response.Body)
// 	byteCount, _ := responseString.Write(content)

// 	fmt.Println(string(content))
// 	fmt.Println("Bite count is:", byteCount)
// 	fmt.Println(responseString.String())
// }

// Post functionality

func PerformPostJsonRequest(myUrl string) {

	// Fake json payload

	requestBody := strings.NewReader(`
		{
			"courseName":"Lets Golang",
			"price": 0,
			"platform" : "youtube.com"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

}

func PerformFormRequest(myUrl string) {

	// formdata

	data := url.Values{}
	data.Add("firstName", "suhail")
	data.Add("lastName", "rabi")
	data.Add("email", "suhail@go.dev")
	data.Add("phone", "123456798")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
