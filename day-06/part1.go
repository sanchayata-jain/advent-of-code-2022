package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := "./input_file.txt"
	dataStream, err := readInputFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	marker := findMarker(dataStream, 4) //for part 1
	fmt.Println(marker)
	messageMarker := findMarker(dataStream, 14)
	fmt.Println(messageMarker) //for part 2
}

func findMarker(dataStream string, seqLength int) int {
	var unique bool
	var marker int
	for i := 0; i < len(dataStream)-seqLength; i++ {
		unique = isUnique(dataStream[i : i+seqLength])
		if unique {
			marker = i + seqLength
			break
		}
	}

	return marker
}

func isUnique(sequence string) bool {
	characters := map[rune]bool{}
	for _, char := range sequence {
		characters[char] = true
	}

	return len(characters) == len(sequence)
}

func readInputFile(filepath string) (string, error) {
	readFile, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	fileScanner := bufio.NewScanner(readFile)
	var line string
	for fileScanner.Scan() {
		line = fileScanner.Text()
	}

	return line, nil
}
