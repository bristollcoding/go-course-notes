package main

import (
	"fmt"
	"io"
	"net/http"
)

func examplemain() {

	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("%w", err)
	}
	io.ReadAll(resp.Body)
	fmt.Println(io.ReadAll(resp.Body))
}
