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
	rounds, err := readInputFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	totalPlayersScore := scoreForRound(rounds)
	fmt.Println(totalPlayersScore) // for part 1

}

// scoreForRound calculates the total score for the player
func scoreForRound(rounds []string) int {
	// loop through rounds
	totalScore := 0 // players score
	for _, round := range rounds {
		playersChoices := strings.Split(round, "")
		opponentChoice := ""
		playersChoice := ""

		switch playersChoices[2] {
		case "X":
			fmt.Println("player picked rock (X)")
			// you picked rock
			totalScore += 1
			playersChoice = "rock"
			opponentChoice = opponentsChoice(playersChoices[0])
			if playersChoice == opponentChoice {
				totalScore += 3
			} else if opponentChoice == "scissors" {
				totalScore += 6
			}
		case "Y":
			// you picked paper
			totalScore += 2
			playersChoice = "paper"
			opponentChoice = opponentsChoice(playersChoices[0])
			if playersChoice == opponentChoice {
				totalScore += 3
			} else if opponentChoice == "rock" {
				totalScore += 6
			}
		case "Z":
			// you picked scissors
			totalScore += 3
			playersChoice = "scissors"
			opponentChoice = opponentsChoice(playersChoices[0])
			if playersChoice == opponentChoice {
				totalScore += 3
			} else if opponentChoice == "paper" {
				totalScore += 6
			}
		}
		

	}

	return totalScore
}

func opponentsChoice(choice string) string {
	opponentChoice := ""
	switch choice {
	case "A":
		// opponent picked rock
		opponentChoice = "rock"
	case "B":
		// opponent picked paper
		opponentChoice = "paper"
	case "C":
		// opponent picker scissors
		opponentChoice = "scissors"
	}
	return opponentChoice
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
