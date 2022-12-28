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

	visibleTrees, _ := partOne(trees)
	fmt.Println(visibleTrees)
}

func partOne(trees [][]string) (int, error) {
	rowLen := len(trees[0])
	colLen := len(trees)
	edgeTrees := ((rowLen*2) + (colLen*2)) - 4
	
	var visibleTrees = 0
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

			maxLeft, err := max(left)
			if err != nil {
				return 0, err
			}
			maxRight, err := max(right)
			if err != nil {
				return 0, err
			}
			maxUp, err := max(up)
			if err != nil {
				return 0, err
			}
			maxDown, err := max(down)
			if err != nil {
				return 0, err
			}

			treeInt, err := strconv.Atoi(tree)
			if err != nil {
				return 0, err
			}

			if treeInt > maxLeft || treeInt > maxRight  || treeInt > maxUp || treeInt > maxDown {
				visibleTrees += 1
			}
		}
		
	}

	return visibleTrees + edgeTrees, nil
}

// max returns the maximum number in a slice
func max(section []string) (int, error) {
	maxNum, err := strconv.Atoi(section[0])
	if err != nil {
		return 0, err
	}
	
	for i := 1; i < len(section); i++ {
		currNumber, err := strconv.Atoi(section[i])
		if err != nil {
			return 0, err
		}
		if currNumber > maxNum {
			maxNum = currNumber
		}
	}

	return maxNum, nil
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
