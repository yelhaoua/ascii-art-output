package asciiart

import (
	"fmt"
	"strings"
)

// func SaveFile(text , name string){

// }

func Output(str string, array [][]string) {
	// out := ""
	// start looping to8 the hieght of line
	for i := 0; i < 8; i++ {
		for _, fin := range str {
			// skip the unprintibal chachacters
			if fin > 126 || fin < 32 {
				continue
			}
			// out += array[int(rune(fin)-32)][i]
			fmt.Print(array[int(rune(fin)-32)][i])
		}
		// out += "\n"
		fmt.Println()
	}
	// SaveFile(out, name)
	// out = ""
}

func PrintSymbole(array [][]string, word string) {
	// check if the word not empty and the array of banner
	if word == "" || len(array) == 0 {
		return
	}
	// change the \n as string to newline
	str := strings.ReplaceAll(word, `\n`, "\n")

	// skip addional newline in print
	if strings.Trim(str, "\n") == "" {
		fmt.Print(str)
		return
	}
	// spliting the word
	words := strings.Split(word, "\n")

	// Print each word 
	for _, s := range words {
		Output(s, array)
	}

}