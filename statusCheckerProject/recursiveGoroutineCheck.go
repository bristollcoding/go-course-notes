package main

import (
	"fmt"
	"net/http"
	"time"
)

func recursiveMain() {
	webs := []string{"http://google.com", "http://facebook.com", "http://golang.org", "http://amazon.com"}
	//create a string channel for comunication between main routine and goroutines
	c := make(chan string)
	for _, web := range webs {
		//call eeach website inside a neew goroutine(concurrent!)
		//pass chan so can be used inside func checkStatus
		go checkStatus(web, c)
		//this way it print all executions but waiting fo message is blocking, so we are back to sequential
		//fmt.Println(<-c)
	}
	//print whatever is send to the channel, as soon as first message appear main routine ends! because it only wait once
	//fmt.Println(<-c)

	//this works like a while true so it runs forever
	for {
		//Launch a new goroutine to check again using web received from previous executed gotroutine
		go checkStatusRec(<-c, c)
	}

	//alternative syntax ( I find it less straightforward...)
	// for w := range c {
	// 	//Launch a new goroutine to check again using web received from previous executed gotroutine
	// 	go checkStatus(w, c)
	// }
}

func checkStatusRec(web string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(web)
	if err != nil {
		fmt.Println(web + " return an error")
		c <- web
	} else {
		fmt.Println(web + " is up and running")
		c <- web
	}
}
