package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := "./input_file.txt"
	input, err := readInputFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fullyContainedPairs, err := calcNumOfFullyContainedPairs(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fullyContainedPairs)
}

func calcNumOfFullyContainedPairs(input []string) (int, error) {
	fullyContainedPairs := 0
	for _, line := range input {
		splitLine := splitIntoNumbers(line)

		lower1, err := strconv.Atoi(splitLine[0])
		if err != nil {
			return 0, err
		}

		upper1, err := strconv.Atoi(splitLine[1])
		if err != nil {
			return 0, err
		}

		lower2, err := strconv.Atoi(splitLine[2])
		if err != nil {
			return 0, err
		}

		upper2, err := strconv.Atoi(splitLine[3])
		if err != nil {
			return 0, err
		}

		if (lower2 >= lower1 && upper1 >= upper2) || (lower1 >= lower2 && upper1 <= upper2) {
			fullyContainedPairs += 1
		}
	}

	return fullyContainedPairs, nil
}

// splitPairs splits the lines from the input file into a slice of the numbers
func splitIntoNumbers(line string) []string {
	pairs := strings.FieldsFunc(line, split)

	return pairs
}

func split(r rune) bool {
	return r == ',' || r == '-'
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
