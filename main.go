package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	printLength := false
	printWordCount := false
	printLineCount := false

	if len(os.Args) == 1 {
		log.Fatal("No file specified")
	}

	if len(os.Args) != 2 {
		for i := 1; i < len(os.Args)-1; i++ {
			if os.Args[i] == "-l" {
				printWordCount = true
			} else if os.Args[i] == "-c" {
				printLength = true
			} else if os.Args[i] == "-w" {
				printWordCount = true
			}
		}
	} else {
		printLength = true
		printWordCount = true
		printLineCount = true
	}

	filePath := os.Args[len(os.Args)-1]

	// Read File
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("The file name should be the last argument")
		log.Fatal(err)
	}

	// Convert to string
	stringContent := string(content)

	// Get all lines
	stringSlice := strings.Split(stringContent, "\n")
	noOfLines := len(stringSlice)

	// Count no of words
	wordCount := 0
	for _, lines := range stringSlice {
		wordCount += len(strings.Split(lines, " "))
	}

	var outputText string
	var output string
	if printLineCount {
		outputText += "\t" + "Lines"
		output += "\t" + fmt.Sprint(noOfLines)
	}

	if printWordCount {
		outputText += "\t" + "Words"
		output += "\t" + fmt.Sprint(wordCount)
	}

	if printLength {
		outputText += "\t" + "Bytes"
		output += "\t" + fmt.Sprint(len(content))
	}

	outputText += "\t" + "File"
	output += "\t" + filePath
	fmt.Println(outputText)
	fmt.Println(output)
}
