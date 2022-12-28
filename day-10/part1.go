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
	instructions, err := readInputFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	signalSum := partOne(instructions)
	fmt.Println(signalSum)
}

func partOne(instructions []string) int {
	signalSum := runInstructions(instructions)
	return signalSum
}

// runInstructions loops through all the instructions in the input file using a for loop
func runInstructions(instructions []string) int {
	x := 1
	i := 1
	signal := 0
	for _, instruction := range instructions {
		if instruction == "noop" {
			if i == 20 || i == 60 || i == 100 || i == 140 || i == 180 || i == 220 {
				signal += calcSignalStrength(i, x)
			}
			i += 1
			continue
		}

		for j := 0; j < 2; j++ {
			if i == 20 || i == 60 || i == 100 || i == 140 || i == 180 || i == 220 {
				signal += calcSignalStrength(i, x)
			}

			
			if j == 0 {
				i += 1
				continue
			}
			
			
			// fmt.Println(i, x)
			x = calcX(x, instruction)
			i += 1
		}
	}

	return signal
}

func calcX(currentX int, instruction string) int {
	instructionSplit := strings.Split(instruction, " ")
	number, _ := strconv.Atoi(instructionSplit[1])

	return currentX + number
}

func calcSignalStrength(cycle int, x int) int {
	return cycle * x
}

func readInputFile(filepath string) ([]string, error) {
	inputFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(inputFile)
	var instructions []string
	for fileScanner.Scan() {
		instruction := fileScanner.Text()
		instructions = append(instructions, instruction)
	}
	return instructions, nil
}
