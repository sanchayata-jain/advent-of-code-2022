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
	numOfPositions := part1(instructions)
	fmt.Println(numOfPositions)
}

func part1(instructions [][]string) int {
	// var sPos = [2]int{0, 0}
	var hPos = [2]int{0, 0}
	var tPos = [2]int{0, 0}
	var tailPositions [][2]int
	tailPositions = append(tailPositions, [2]int{0, 0})
	for _, instruction := range instructions {
		direction := instruction[0]
		distanceInt, _ := strconv.Atoi(instruction[1])
		for i := 0; i < distanceInt; i++ {
			// fmt.Println(hPos)
			hPos = moveHead(hPos, direction, 1)
			tailTouching := isTailTouching(tPos, hPos) //tailTouching is either true or false,
			// fmt.Println(tPos, hPos)
			if !tailTouching {
				// need to move tail to be touching head
				// fmt.Println(tPos, hPos)
				// fmt.Println("tail is not touching")
				fmt.Println(tPos)
				tPos = moveTail(tPos, hPos)
				fmt.Println(tPos)
				// fmt.Println()

				tPosExists := checkTailPositionExists(tPos, tailPositions)
				fmt.Println(tPosExists)
				fmt.Println()
				if !tPosExists {
					tailPositions = append(tailPositions, tPos)
				}
			}
		}
	}
	// fmt.Println(tailPositions)
	return len(tailPositions)
}

func checkTailPositionExists(tPos [2]int, tailPositions [][2]int) bool {
	for _, position := range tailPositions {
		if tPos == position {
			return true
		}
	}
	return false
}

func moveTail(tPos [2]int, hPos [2]int) [2]int {
	// up and right positive
	// down and left negative
	var hDistance int
	var vDistance int
	horizontalDiff := hPos[1] - tPos[1]
	verticalDiff := hPos[0] - tPos[0]
	if horizontalDiff == 2 && verticalDiff == 0 {
		hDistance = 1
	} else if horizontalDiff == -2 && verticalDiff == 0 {
		hDistance = -1
	}
	if verticalDiff == 2 && horizontalDiff == 0 {
		vDistance = 1
	} else if vDistance == -2 && horizontalDiff == 0 {
		vDistance = -1
	}

	if horizontalDiff != 0 && verticalDiff != 0 {
		if horizontalDiff > 0 && verticalDiff > 0 {
			// move tail diagonally up and right
			hDistance = 1
			vDistance = 1
		}
		if horizontalDiff < 0 && verticalDiff < 0 {
			// move tail diagonally down and left
			hDistance = -1
			vDistance = -1
		}
		if horizontalDiff > 0 && verticalDiff < 0 {
			// move tail diagonally right and down
			hDistance = 1
			vDistance = -1
		}
		if horizontalDiff < 0 && verticalDiff > 0 {
			// move tail diagonally left and up
			hDistance = -1
			vDistance = 1
		}
	}
	newTailPosVertical := tPos[0] + vDistance
	newTailPosHorizontal := tPos[1] + hDistance

	return [2]int{newTailPosVertical, newTailPosHorizontal}
}

// isTailTouching returns true or false.
// tail is touching when tail is next to head either vertically or horizontally, overlapping, or diagonally adjacent
func isTailTouching(tPos [2]int, hPos [2]int) bool {
	// case 1: head is one space right to tail
	if hPos[1] == tPos[1]+1 && hPos[0] == tPos[0] || 
	hPos[1] == tPos[1]-1 && hPos[0] == tPos[0] ||
	hPos[0] == tPos[0]+1 && hPos[1] == tPos[1] || 
	hPos[0] == tPos[0]-1 && hPos[1] == tPos[1] ||
	hPos[0] == tPos[0]+1 && hPos[1] == tPos[1]+1 || 
	hPos[0] == tPos[0]-1 && hPos[1] == tPos[1]-1 ||
	hPos[0] == tPos[0]-1 && hPos[1] == tPos[1]+1 || 
	hPos[0] == tPos[0]+1 && hPos[1] == tPos[1]-1 ||
	hPos[0] == tPos[0] && hPos[1] == tPos[1] {
		return true
	}

	return false
}

func moveHead(headPos [2]int, direction string, distance int) [2]int {
	// up and right positive
	// down and left negative
	var newHeadPos = [2]int{0, 0}
	switch direction {
	case "U":
		// up
		pos := headPos[0] + distance
		newHeadPos = [2]int{pos, headPos[1]}
	case "R":
		// right
		pos := headPos[1] + distance
		newHeadPos = [2]int{headPos[0], pos}
	case "D":
		// down
		pos := headPos[0] - distance
		newHeadPos = [2]int{pos, headPos[1]}
	case "L":
		// left
		pos := headPos[1] - distance
		newHeadPos = [2]int{headPos[0], pos}
	}
	return newHeadPos
}

func readInputFile(filepath string) ([][]string, error) {
	readFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	var instructions [][]string
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		instruction := fileScanner.Text()
		instructionSplit := strings.Split(instruction, " ")
		instructions = append(instructions, instructionSplit)
	}
	return instructions, nil
}
