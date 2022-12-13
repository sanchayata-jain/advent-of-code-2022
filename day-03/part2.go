package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type commonItems map[rune]bool

func main() {
	filepath := "./input_file1.txt"
	inputFile, err := readInputFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	totalPriority := findCommonItem(inputFile)
	fmt.Println(totalPriority)

}

func findCommonItem(input []string) int {
	priority := 0
	for _, line := range input {
		items := commonItems{}
		second := false
		for _, letter := range line {
			if string(letter) == " " {
				second = true
				continue
			}
			if second {
				// stop adding to map and just check if letters now exist in current map
				if _, exists := items[letter]; exists {
					items[letter] = true
				}
				if string(letter) == " " {
					third = true
				}
					// check third line to see if letters exist in current map
					if v, exists := items[letter]; exists {
						if v {
							// was in first line, second line and now third line so we can convert
							priority += convert(letter)
							break
						}
					}
				} else {
					continue
				}
			}
			items[letter] = false
		}
	}
	return priority
}

func convert(letter rune) int {
	numVal := 0
	if letter >= 'a' && letter <= 'z' {
		numVal = int(letter - 'a' + 1)
	}
	if letter >= 'A' && letter <= 'Z' {
		numVal = int(letter - 'A' + 27)
	}
	return numVal
}

func readInputFile(filepath string) ([]string, error) {
	readFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	block := ""
	i := 0
	for fileScanner.Scan() {
		block += fileScanner.Text() + " "
		if (i+1)%3 == 0 {
			fileLines = append(fileLines, block)
		}
		i += 1
	}

	readFile.Close()

	return fileLines, nil
}
