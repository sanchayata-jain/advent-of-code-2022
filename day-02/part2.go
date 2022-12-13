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
	inputFile, err := readInputFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	totalScore := calcPlayersChoice(inputFile)
	fmt.Println(totalScore)
}

// calcPlayersChoice figures out what the player should choose depending on the opponents choice and 2nd letter in input file
func calcPlayersChoice(input []string) int {
	playersChoice := ""
	score := 0
	for _, round := range input {
		choices := strings.Split(round, "")
		switch choices[2] {
		case "Y":
			// player needs to be a draw
			playersChoice = choices[0] //choices[0] is the opponents choice
			score += calcScore(playersChoice) + 3 

		case "X":
			// player needs to lose
			if choices[0] == "A" {
				// opponent chose rock
				playersChoice = "scissors"
			} else if choices[0] == "B" {
				//opponent chose paper
				playersChoice = "rock"
			} else {
				//opponent chose scissors
				playersChoice = "paper"
			}
			score += calcScore(playersChoice)

		case "Z":
			// player needs to win
			if choices[0] == "A" {
				playersChoice = "paper"
			} else if choices[0] == "B" {
				playersChoice = "scissors"
			} else {
				playersChoice = "rock"
			}
			score += calcScore(playersChoice) + 6
		}

	}
	return score
}

func calcScore(pc string) int {
	score := 0
	switch pc {
	case "A":
		score += 1
	case "B":
		score += 2
	case "C":
		score += 3
	case "rock":
		score += 1
	case "paper":
		score += 2
	case "scissors":
		score += 3
	}

	return score
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
