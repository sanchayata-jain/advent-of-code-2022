package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filepath := "./input_file.txt"
	drawing, _, err := readInputFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(drawing)
}

func read(filepath string) {
	readFile, err := os.Open(filepath)
	if err != nil {
		// return soemthing
	}
	fileScanner := bufio.NewScanner(readFile)
	// fileScanner.Split()
}

// readInputFile will return both the drawing and the procedure so they are seperated from each other
func readInputFile(filepath string) ([][]string, []string, error) {
	readFile, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var drawing [][]string
	var procedure []string
	var stack1 []string
	var stack2 []string
	var stack3 []string
	var stack4 []string
	var stack5 []string
	var stack6 []string
	var stack7 []string
	var stack8 []string
	var stack9 []string
	stillDrawing := true
	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), "1") || string(fileScanner.Text()) == "" {
			//entering the procedure part of the file
			stillDrawing = false
			continue
		}
		if stillDrawing {
			// lineSplit := strings.Split(fileScanner.Text(), " ")
			stack1 = append(stack1, string(fileScanner.Text()[1]))
			stack2 = append(stack2, string(fileScanner.Text()[5]))
			stack3 = append(stack3, string(fileScanner.Text()[9]))
			stack4 = append(stack4, string(fileScanner.Text()[13]))
			stack5 = append(stack5, string(fileScanner.Text()[17]))
			stack6 = append(stack6, string(fileScanner.Text()[21]))
			stack7 = append(stack7, string(fileScanner.Text()[25]))
			stack8 = append(stack8, string(fileScanner.Text()[29]))
			stack9 = append(stack9, string(fileScanner.Text()[33]))
		}
		if !stillDrawing {
			procedure = append(procedure, fileScanner.Text())
		}
	}

	drawing = append(drawing, stack1, stack2, stack3, stack4, stack5, stack6, stack7, stack8, stack9)

	// fmt.Println()


	readFile.Close()

	return drawing, procedure, nil
}
