package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items              []int
	operation          []string
	test               int
	testTrueCondition  int
	testFalseCondition int
}

func main() {
	filepath := "./test_file.txt"
	scanner, err := readInputFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	monkeysInfo, err := parse(scanner)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range monkeysInfo {
		fmt.Println(info)
	}
}

// grab first starting item
// do operation with starting item, this results in worry level
// monkey gets bored -> divide worry level by 3, new result is rounded down to new worry level
// take new worry score and see if its divisible by number stored in test
// if true -> do true condition
// if false -> do false condition

func parse(scanner *bufio.Scanner) ([]Monkey, error) {
	var monkeys []Monkey
	monkey := Monkey{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		// startingItemsInt := []int{}
		operation := []string{}
		line := strings.Split(scanner.Text(), ": ")
		if strings.Contains(line[0], "Starting items") {
			startingItems := strings.Split(strings.ReplaceAll(line[1], " ", ""), ",")
			for _, startingItem := range startingItems {
				startingItemInt, err := strconv.Atoi(startingItem)
				if err != nil {
					return nil, err
				}
				// startingItemsInt = append(startingItemsInt, startingItemInt)
				monkey.items = append(monkey.items, startingItemInt)
			}
		}
		if strings.Contains(line[0], "Operation") {
			operation = strings.Split(line[1], " ")
			monkey.operation = operation
		}
		if strings.Contains(line[0], "Test") {
			testSplit := strings.Split(line[1], " ")
			divisibleNum, err := strconv.Atoi(testSplit[len(testSplit)-1])
			if err != nil {
				return nil, err
			}
			monkey.test = divisibleNum
		}
		if strings.Contains(line[0], "true") {
			lineSplit := strings.Split(line[1], " ")
			monkeyNum, err := strconv.Atoi(lineSplit[len(lineSplit)-1])
			if err != nil {
				return nil, err
			}
			monkey.testTrueCondition = monkeyNum
		}
		if strings.Contains(line[0], "false") {
			lineSplit := strings.Split(line[1], " ")
			monkeyNum, err := strconv.Atoi(lineSplit[len(lineSplit)-1])
			if err != nil {
				return nil, err
			}
			monkey.testFalseCondition = monkeyNum
			monkeys = append(monkeys, monkey)
			monkey = Monkey{}
		}
	}
	return monkeys, nil
}

func readInputFile(filepath string) (*bufio.Scanner, error) {
	inputFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(inputFile)

	return scanner, nil
}
