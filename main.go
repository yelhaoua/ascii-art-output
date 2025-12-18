package main

import (
	"os"

	asciiart "asciiart/func"
)

func main() {
	// check the num of args
	if len(os.Args) > 4 || len(os.Args) == 1 {
		return
	}

	Content := asciiart.Splite("standard")
	asciiart.PrintSymbole(Content, os.Args[1])

}
