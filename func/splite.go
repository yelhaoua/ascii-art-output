package asciiart

import (
	"os"
	"strings"
)

func Splite(fileName string) [][]string {
	var symbole []string
	var allData [][]string
	
	// read the bannere from banners folder
	data, err := os.ReadFile("./banners/" + fileName + ".txt")
	if err != nil {
		return [][]string{}
	}

	// splite the content of bannere whit "\n"
	symbole = strings.Split(string(data), "\n")

	// add each symbole the an array
	for i := 1; i < len(symbole); i += 9 {
		if i+8 < len(symbole) {
			allData = append(allData, symbole[i:i+8])
		}
	}
	return allData
}
