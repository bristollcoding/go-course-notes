package main

import (
	"fmt"
	"net/http"
)

func basicgoroutineMain() {
	webs := []string{"http://google.com", "http://facebook.com", "http://golang.org", "http://amazon.com"}
	//create a string channel for comunication between main routine and goroutines
	c := make(chan string)
	for _, web := range webs {
		//call eeach website inside a neew goroutine(concurrent!)
		//pass chan so can be used inside func checkStatus
		go checkStatusOnce(web, c)
		//this way it print all executions but waiting fo message is blocking, so we are back to sequential
		//fmt.Println(<-c)
	}
	//print whatever is send to the channel, as soon as first message appear main routine ends! because it only wait once
	//fmt.Println(<-c)

	//this way after all goroutines are lanunched main routine will wait for as many messages as nedded (1 message for check in this case)
	for i := 0; i < len(webs); i++ {
		fmt.Println(<-c)
	}
}

func checkStatusOnce(web string, c chan string) {
	_, err := http.Get(web)
	if err != nil {
		fmt.Println(web + " return an error")
		c <- "down"
	} else {
		fmt.Println(web + " is up and running")
		c <- "up"
	}
}
