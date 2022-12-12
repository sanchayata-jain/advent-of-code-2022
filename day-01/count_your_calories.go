package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filepath := "../input_file.txt"
	caloriesList, err := readInputFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	elfNum, err := calorieCounter(caloriesList)
	if err != nil {
		log.Printf("calorie Counter: %s", err)
	}
	fmt.Print(elfNum)
}

// calorieCounter sums the number of calories each elf consumes and returns the highest sum of calories consumed by an elf.
func calorieCounter(caloriesList []string) (int, error) {
	caloriesList = append(caloriesList, "")
	maxCalorieSum := 0
	sum := 0
	for _, calorie := range caloriesList {
		if calorie == "" {
			if sum > maxCalorieSum {
				maxCalorieSum = sum
			}
			sum = 0
			continue
		}
		calorieInt, err := strconv.Atoi(calorie)
		if err != nil {
			return 0, err
		}
		sum += calorieInt
	}

	return maxCalorieSum, nil
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
