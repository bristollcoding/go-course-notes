package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//Read argument from terminal
	args := os.Args[1:]
	if len(args) > 0 {

		fmt.Println(args[0])
		//open file with given filename (fn)
		fn := args[0]
		f, err := os.Open(fn)
		if err != nil {
			fmt.Printf("Error opening file: %s wtih error: %v", fn, err)
			os.Exit(1)
		}
		//use cpy to pass stream from fileReader to terminalWriter
		io.Copy(os.Stdout, f)
	} else {
		fmt.Println("Filename arg missing")
	}

}
