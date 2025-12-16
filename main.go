package main

import (
	"fmt"
	"os"

	asciiart "asciiart/func"
)

func main() {
	// check the num of args
	if len(os.Args) > 4 || len(os.Args) == 1 {
		return
	}
	// add the logic for run program in all cases just temp
	switch len(os.Args) {
	case 2:
		Content := asciiart.Splite("standard")
		asciiart.PrintSymbole(Content, os.Args[1])
	case 3:
		Content := asciiart.Splite(os.Args[2])
		asciiart.PrintSymbole(Content, os.Args[1])
	case 4:
		fmt.Println(os.Args[1][0:8])
		if os.Args[1][0:9] != "--output=" {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			return
		}
		// Split the banner (os.Args[2] == name of banner)
		Content := asciiart.Splite(os.Args[3])
		// Print the text of user input  (os.Args[1] == text)
		asciiart.PrintSymbole(Content, os.Args[2])
	}
}
