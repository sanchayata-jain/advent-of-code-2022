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
	fullyContainedPairs, overlappingPairs, err := calcNumOfPairs(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fullyContainedPairs) // for part 1
	fmt.Println(overlappingPairs) // for part 2
}

// this function returns number of fully contained pairs and number of overlapping pairs (which contains fully contained as well)
func calcNumOfPairs(input []string) (int, int, error) {
	fullyContainedPairs := 0
	overlap := 0
	for _, line := range input {
		lower1, upper1, lower2, upper2, err := splitIntoNumbers(line)
		if err != nil {
			return 0, 0, err
		}

		// fully contained pairs
		if (lower2 >= lower1 && upper1 >= upper2) || (lower1 >= lower2 && upper1 <= upper2) {
			fullyContainedPairs += 1
		}

		// overlapping pairs
		if (lower2 >= lower1 && lower2 <= upper1 || lower1 <= upper2 && lower1 >= lower2) {
			overlap += 1
		}
	}

	return fullyContainedPairs, overlap, nil
}

// splitPairs splits the lines from the input file into a slice of the numbers
func splitIntoNumbers(line string) (int, int, int, int, error) {
	numbers := strings.FieldsFunc(line, split)

	lower1, err := strconv.Atoi(numbers[0])
		if err != nil {
			return 0, 0, 0, 0, err
		}

	upper1, err := strconv.Atoi(numbers[1])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	lower2, err := strconv.Atoi(numbers[2])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	upper2, err := strconv.Atoi(numbers[3])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return lower1, upper1, lower2, upper2, nil
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
