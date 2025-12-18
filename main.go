package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "asciiart/func"
)

func main() {
	// check the num of args
	if len(os.Args) > 4 || len(os.Args) == 1 {
		return
	}
	if len(os.Args) == 2 {
		if os.Args[1] == "" {
			return
		}
		Content := asciiart.Splite("standard")
		fmt.Print(asciiart.PrintSymbole(Content, os.Args[1]))
	} else if len(os.Args) == 3 {
		if os.Args[1] == "" || os.Args[2] == "" {
			return
		}
		Content := asciiart.Splite(os.Args[2])
		fmt.Print(asciiart.PrintSymbole(Content, os.Args[1]))
	} else if len(os.Args) == 4 {
		if strings.HasSuffix(os.Args[1], ".txt") && strings.HasPrefix(os.Args[1], "--output=") && os.Args[2] != "" {
			if os.Args[3] == "" {
				asciiart.PrintInFile(os.Args[1], os.Args[2], "standard")
			} else {
				asciiart.PrintInFile(os.Args[1], os.Args[2], os.Args[3])
			}
		} else {
			printErr()
		}
	}
}

func printErr() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run .  something ")
}
