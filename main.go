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
			printErr(1)
			return
		}
		Content := asciiart.Splite("standard")
		asciiart.PrintSymbole(Content, os.Args[1])
	} else if len(os.Args) == 3 {
		if os.Args[2] == "" || os.Args[3] == "" {
			printErr(2)
			return
		}
		Content := asciiart.Splite(os.Args[2])
		asciiart.PrintSymbole(Content, os.Args[1])
	} else if len(os.Args) == 4 {
		if strings.HasSuffix(os.Args[1], ".txt") && strings.HasPrefix(os.Args[1], "--output=") && os.Args[2] != "" || os.Args[3] != "" {
			asciiart.PrintInFile(os.Args[1], os.Args[2], os.Args[3])
		} else {
			printErr(3)
		}
	}

}

func printErr(option int) {
	switch option {
	case 1:
		fmt.Println("Usage: go run . [STRING] ")
		fmt.Println()
		fmt.Println("EX: go run .  something ")
	case 2:
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("EX: go run . something standard")
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
	}
}
