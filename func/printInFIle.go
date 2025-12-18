package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func PrintInFile(outputName string, word string, fileName string) {
	// Split the banner
	array := Splite(fileName)
	// get the string as ascii art
	str := PrintSymbole(array, word)
	// create the file  of output
	file, err := os.Create(outputName[strings.Index(outputName, "=")+1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	// write the string in the output
	_, err = file.WriteString(str)
	if err != nil {
		fmt.Println("Error to Write the ascii art")
		return
	}

}
