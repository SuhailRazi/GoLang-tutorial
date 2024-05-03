package main

import (
	"fmt"
	"io"

	"net/http"
)

const url = "https://www.youtube.com/"

func main() {
	fmt.Println("Lco web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type : %T \n", response)

	// response.Body.Close() // Callers responsibility

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println(content)

}
