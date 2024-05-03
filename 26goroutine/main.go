package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

func main() {
	// go greeter("hello")
	// greeter("World")
	websiteList := []string{
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("signals", signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endPoint string) {
	defer wg.Done()
	res, err := http.Get(endPoint)

	if err != nil {
		panic(err)
	} else {
		mut.Lock()
		signals = append(signals, endPoint)
		mut.Unlock()
		fmt.Printf("%d status code for website %s \n", res.StatusCode, endPoint)
	}
}
