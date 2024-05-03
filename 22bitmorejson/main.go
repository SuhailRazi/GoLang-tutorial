package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json video")
	encodeJson()
	decodeJson()
}

func encodeJson() {
	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "Learcode online", "abc123", []string{"webdev", "js"}},
		{"Nodejs Bootcamp", 199, "Youtube online", "gfh458", []string{"backend", "express"}},
		{"GOlang Bootcamp", 299, "W2school online", "etyr5", nil},
	}

	// Package this data as json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJson)
}

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs Bootcamp",
		"Price": 299,
		"website": "Learcode online",
		"Tags": ["webdev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSOn was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	// some cases where you want to add data to key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v \n", myOnlineData)

	for a, b := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T \n", a, b, b)
	}

}
