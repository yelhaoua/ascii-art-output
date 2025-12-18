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
	if len(os.Args) == 2 {
		Content := asciiart.Splite("standard")
		asciiart.PrintSymbole(Content, os.Args[1])
	} else if len(os.Args) == 3 {
		Content := asciiart.Splite(os.Args[2])
		asciiart.PrintSymbole(Content, os.Args[1])
	} else if len(os.Args) == 4 {
		asciiart.PrintInFile(os.Args[1], os.Args[2], os.Args[3])
	}

}
