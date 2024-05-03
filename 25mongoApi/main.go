package main

import (
	"fmt"
	"log"
	"mongoApi/router"
	"net/http"
)

func main() {
	fmt.Println("Mongo db api")
	r := router.Router()
	fmt.Println("server is getting started..")

	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening at port 400")
}
