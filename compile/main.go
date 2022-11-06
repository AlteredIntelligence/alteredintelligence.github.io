package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	files, _ := os.ReadDir("../img")
	thisList := []string{}
	for _, s := range files {
		if !strings.Contains(s.Name(), ".DS_Store") {
			thisList = append(thisList, "<img src=\"img/"+s.Name()+"\" class=\"img-fluid rounded\">\n")
		}

	}
	reverseList := reverse(thisList)

	fmt.Println(reverseList)
}

func reverse(input []string) []string {
	var output []string

	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}
