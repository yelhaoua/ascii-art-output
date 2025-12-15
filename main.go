package main

import (
	asciiart "asciiart/func"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	} else {
		file := asciiart.Splite(os.Args[1])
		asciiart.PrintSymbole(file, os.Args[2])
	}

}
