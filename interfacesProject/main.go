package main

import (
	"fmt"
	"io"
	"net/http"
)

// customWriter
type customWriter struct{}

func writerMain() {

	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("%w", err)
	}
	//direct use of reader interface
	//var respBytes []byte
	//Read function doesn't resize byte slice!!
	// respBytes := make([]byte, 99999)
	// nb, err2 := resp.Body.Read(respBytes)
	// if err != nil {
	// 	fmt.Println(err2)
	// }
	//fmt.Printf("read bytes: %v, \n body: %v", nb, string(respBytes))
	//using helpers
	//io.Copy(os.Stdout, resp.Body)

	//using our custom writer
	cr := customWriter{}
	io.Copy(cr, resp.Body)
}

// this make customWriter implement writer interface
func (customWriter) Write(bs []byte) (int, error) {
	fmt.Println("CustomWriter:")
	fmt.Println(string(bs))
	fmt.Println("customWriter end")
	return len(bs), nil
}
