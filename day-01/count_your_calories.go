package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	filepath := "../input_file.txt"
	caloriesList, err := readInputFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	highestCaloriesSum, calorieSumList, err := calorieCounter(caloriesList)
	if err != nil {
		log.Printf("calorie Counter: %s", err)
	}
	sum := topThreeCaloriesSum(calorieSumList)

	fmt.Println(highestCaloriesSum) // for part 1
	fmt.Println(sum) // for part 2
}

// calorieCounter sums the number of calories each elf consumes and returns the highest sum of calories consumed by an elf
// and a list of the sum of calories
func calorieCounter(caloriesList []string) (int, []int, error) {
	caloriesList = append(caloriesList, "")
	var caloriesSumList []int
	maxCalorieSum := 0
	sum := 0
	for _, calorie := range caloriesList {
		if calorie == "" {
			caloriesSumList = append(caloriesSumList, sum)
			if sum > maxCalorieSum {
				maxCalorieSum = sum
			}
			sum = 0
			continue
		}
		calorieInt, err := strconv.Atoi(calorie)
		if err != nil {
			return 0, nil, err
		}
		sum += calorieInt
	}

	return maxCalorieSum, caloriesSumList, nil
}

// topThreeCaloriesSum sorts the caloriesSumList and returns the sum of the top three highest calories sum
func topThreeCaloriesSum(caloriesSumList []int) int{
	sort.Ints(caloriesSumList)
	sum := 0
	for i := 1; i <= 3; i++ {
		sum += caloriesSumList[len(caloriesSumList) - i]
	}
	return sum
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
