package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	printLength := false
	printWordCount := false
	printLineCount := false
	printFile := true

	// if len(os.Args) == 1 {
	// 	log.Fatal("No file specified")
	// }
	var stringContent string
	var count int
	noStdin := false

	filePath := os.Args[len(os.Args)-1]
	if filePath == "-l" || filePath == "-c" || filePath == "-w" || filePath == os.Args[0] {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stringContent += scanner.Text() + "\n"
		}
		printFile = false
		count = len(os.Args)
	} else {
		// Read File
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("The file name should be the last argument")
			log.Fatal(err)
		}
		stringContent = string(content)
		count = len(os.Args) - 1
		noStdin = true
	}

	if len(os.Args) == 1 || (noStdin && len(os.Args) == 2) {
		printLength = true
		printWordCount = true
		printLineCount = true
	} else {
		for i := 1; i < count; i++ {
			if os.Args[i] == "-l" {
				printLineCount = true
			} else if os.Args[i] == "-c" {
				printLength = true
			} else if os.Args[i] == "-w" {
				printWordCount = true
			} else {
				log.Fatal("Invalid flag, only -l -c -w are supported")
			}
		}
	}

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
		output += "\t" + fmt.Sprint(len(stringContent))
	}

	if printFile {
		outputText += "\t" + "File"
		output += "\t" + filePath
	}

	fmt.Println(outputText)
	fmt.Println(output)
}
