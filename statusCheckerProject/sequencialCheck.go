package main

import (
	"fmt"
	"net/http"
)

func sequentialMain() {
	webs := []string{"http://google.com", "http://facebook.com", "http://golang.org", "http://amazon.com"}
	for _, web := range webs {
		//call website
		checkStatusSeq(web)
	}
}

func checkStatusSeq(web string) {
	_, err := http.Get(web)
	if err != nil {
		fmt.Println(web + " return an error")

	} else {
		fmt.Println(web + " is up and running")
	}
}
