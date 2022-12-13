package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type commonCompartmentItems map[rune]struct{}

func main() {
	filepath := "./input_file.txt"
	input, err := readInputFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	totalPriority := findCommonItem(input)
	fmt.Println(totalPriority)

}

func findCommonItem(input []string) int {
	priority := 0
	for _, rucksack := range input {
		items := commonCompartmentItems{}
		type void struct{}
		var voidItem void
		first, second := splitRucksackIntoCompartments(rucksack)
		for _, letter := range first {
			items[letter] = voidItem
		}
		for _, letter2 := range second {
			if _, exists := items[letter2]; exists {
				priority += convert(letter2)
				break
			}
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

// splits the string in half and returns two strings representing the two rucksack compartments
func splitRucksackIntoCompartments(rucksack string) (string, string) {
	halfLen := len(rucksack) / 2
	firstCompartment := rucksack[0:halfLen]
	secondCompartment := rucksack[halfLen:]

	return firstCompartment, secondCompartment
}

func readInputFile(filepath string) ([]string, error) {
	readFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines, nil
}
