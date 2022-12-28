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
	trees, err := readInputFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	visibleTrees, scenic, _ := partOne(trees)
	fmt.Println(visibleTrees)
	fmt.Println(scenic)
}

// func partTwo(trees [][]string) {
// 	for i, row := range trees {
// 		if i == 0 {
// 			// first row so
// 		}
// 	}
// }

func partOne(trees [][]string) (int, int, error) {
	rowLen := len(trees[0])
	colLen := len(trees)
	edgeTrees := ((rowLen * 2) + (colLen * 2)) - 4

	var visibleTrees = 0
	var scenicScore []int
	for i, row := range trees {
		if i == 0 {
			// first row is visible so skip
			continue
		}
		if i == len(trees)-1 {
			// last row is visible so exit loop
			break
		}
		for j, tree := range row {
			visibleLeft := false
			visibleRight := false
			visibleUp := false
			visibleDown := false
			// visibleTree := false
			if j == 0 {
				// first column of each row is visible so skip
				continue
			}
			if j == len(row)-1 {
				// last column of each row is visible so skip to next row
				break
			}

			left := row[:j]    // everything left of current tree in the same row (not including the current tree)
			right := row[j+1:] // everything right of current tree in the same row (not including the current tree)
			var up []string
			var down []string
			for k := 0; k < i; k++ {
				up = append(up, trees[k][j])
			}

			for k := i + 1; k < len(trees); k++ {
				down = append(down, trees[k][j])
			}

			maxLeft, _, err := max(left)
			if err != nil {
				return 0, 0, err
			}
			maxRight, _, err := max(right)
			if err != nil {
				return 0, 0, err
			}
			maxUp, _, err := max(up)
			if err != nil {
				return 0, 0, err
			}
			maxDown, _, err := max(down)
			if err != nil {
				return 0, 0, err
			}

			treeInt, err := strconv.Atoi(tree)
			if err != nil {
				return 0, 0, err
			}

			// for part 1
			if treeInt > maxLeft || treeInt > maxRight || treeInt > maxUp || treeInt > maxDown {
				visibleTrees += 1
				// visibleTree = true
			}
			// fmt.Println(treeInt, maxLeft)
			// fmt.Println(treeInt, maxRight)
			// fmt.Println(treeInt, maxUp)
			// fmt.Println(treeInt, maxDown)

			if treeInt > maxLeft {
				visibleLeft = true
			}
			if treeInt > maxRight {
				visibleRight = true
			}
			if treeInt > maxUp {
				visibleUp = true
			}
			if treeInt > maxDown {
				visibleDown = true
			}

			// fmt.Println(visibleLeft)
			// fmt.Println(visibleRight)
			// fmt.Println(visibleUp)
			// fmt.Println(visibleDown)
			// fmt.Println()

			maxLeftIndex := findIndex(left, treeInt, "left")
			maxRightIndex := findIndex(right, treeInt, "right")
			maxUpIndex := findIndex(up, treeInt, "up")
			maxDownIndex := findIndex(down, treeInt, "down")

			leftDiff := 0
			rightDiff := 0
			upDiff := 0
			downDiff := 0

			//calculate scenic score for part 2

			if !visibleLeft {
				// fmt.Println(maxLeftIndex)
				leftDiff = j - maxLeftIndex
			} else if visibleLeft {
				leftDiff = j
			}
			if !visibleRight {
				rightDiff = maxRightIndex + 1
			} else if visibleRight {
				rightDiff = len(row) - 1 - j
			}
			if !visibleUp {
				upDiff = i - maxUpIndex
			} else if visibleUp {
				upDiff = i
			}
			if !visibleDown {
				downDiff = maxDownIndex + 1
			} else if visibleDown {
				downDiff = colLen - 1 - i
			}
			// fmt.Println(i, j)
			// fmt.Println(visibleLeft)
			// fmt.Println(leftDiff)
			// fmt.Println(visibleRight)
			// fmt.Println(rightDiff)
			// fmt.Println(visibleUp)
			// fmt.Println(upDiff)
			// fmt.Println(visibleDown)
			// fmt.Println(downDiff)
			// fmt.Println()
			scenicScore = append(scenicScore, leftDiff*rightDiff*upDiff*downDiff)
		}

	}

	maxScenicScore := max2(scenicScore)

	return visibleTrees + edgeTrees, maxScenicScore, nil
}

func max2(section []int) int {
	maxNum := section[0]
	for i := 1; i < len(section); i++ {
		currNum := section[i]
		if currNum >= maxNum {
			maxNum = currNum
		}
	}
	return maxNum
}

func findIndex(section []string, treeVal int, direction string) (int) {
	// finds index for number which is bigger than current tree in same row or same col (left, right, up, down)
	closestIdx := 0
	if direction == "left" || direction == "up" {
		for i := range section {
			otherTree, _ := strconv.Atoi(section[i])
			if otherTree >= treeVal {
				closestIdx = i
			}
		}
		return closestIdx
	}
	if direction == "right" || direction == "down" {
		for i := range section {
			otherTree, _ := strconv.Atoi(section[i])
			if otherTree >= treeVal {
				return i
			}
		}
	}
	return 0
}

// max returns the maximum value and index in a slice
func max(section []string) (int, int, error) {
	maxNum, err := strconv.Atoi(section[0])
	if err != nil {
		return 0, 0, err
	}
	maxIndex := 0

	for i := 1; i < len(section); i++ {
		currNumber, err := strconv.Atoi(section[i])
		if err != nil {
			return 0, 0, err
		}
		if currNumber > maxNum {
			maxNum = currNumber
			maxIndex = i
		}
	}

	return maxNum, maxIndex, nil
}

func readInputFile(filepath string) ([][]string, error) {
	readFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	// var line string
	var trees [][]string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		row := strings.Split(line, "")
		trees = append(trees, row)
	}

	return trees, nil
}
