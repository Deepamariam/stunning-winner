package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var scoreUser int = 0
var scoreComputer int = 0
var userInputTitle string
var toContinue string
var toContinueLower string

func main() {

	fmt.Printf("\n\t\tStone Paper Scissor!\n\n")
	stonePaperScissor()
	for {
		playAgain()
	}

}

func stonePaperScissor() {

	for {
		if scoreUser >= 5 || scoreComputer >= 5 {
			if scoreUser == 5 {
				fmt.Printf("\t\t\tYou won the game!\n\n")
			} else {
				fmt.Printf("\t\t\tComputer won the game!\n\n")
			}
			break
		}
		userInput := getInput()
		userInputTitle = strings.Title(userInput)
		if userInputTitle == "Stone" || userInputTitle == "Scissor" || userInputTitle == "Paper" {

			fmt.Printf("\tYou chose %q.\n\n", strings.Title(userInput))
			systemInput := computerInput()
			fmt.Printf("\tComputer chose %q.\n\n", systemInput)
			compareInput(userInputTitle, systemInput)
		} else {
			fmt.Printf("Invalid Choice. Choose Stone, Paper, or Scissor!\n\n")
			continue
		}
	}

}

func getInput() string {

	var userInput string
	fmt.Println("Choose Stone, Paper, or Scissor!")
	fmt.Scanln(&userInput)
	return userInput

}

func computerInput() string {

	choice := []string{"Stone", "Paper", "Scissor"}
	computerInput := choice[rand.Intn(len(choice))]
	return computerInput

}

func compareInput(inputOne, inputTwo string) {

	switch inputOne {
	case "Scissor":
		if inputTwo == "Paper" {
			fmt.Println("\nYay, you scored!")
			scoreUser = scoreUser + 1
		} else if inputTwo == "Stone" {
			fmt.Println("\nOops, you lost!")
			scoreComputer = scoreComputer + 1
		} else {
			fmt.Printf("\nIt's neither a win nor lose!\nno")
		}

	case "Paper":
		if inputTwo == "Stone" {
			fmt.Println("\nYay, you scored!")
			scoreUser = scoreUser + 1
		} else if inputTwo == "Scissor" {
			fmt.Println("\nOops, you lost!")
			scoreComputer = scoreComputer + 1
		} else {
			fmt.Printf("\nIt's neither a win nor lose!")
		}

	case "Stone":
		if inputTwo == "Scissor" {
			fmt.Printf("\nYay, you scored!\n")
			scoreUser = scoreUser + 1
		} else if inputTwo == "Paper" {
			fmt.Printf("\nOops, you lost!\n")
			scoreComputer = scoreComputer + 1
		} else {
			fmt.Printf("\nIt's neither a win nor lose!\n")
		}
	}

	printScore(scoreUser, scoreComputer)
}

func printScore(scoreOne, scoreTwo int) {

	fmt.Printf("\n\nYour Score: %d\n", scoreOne)
	fmt.Printf("\n\nComputer's Score: %d\n\n", scoreTwo)
}

func playAgain() {

	fmt.Println("Do you want to you continue playing? Yes or No!")
	fmt.Scanln(&toContinue)
	toContinueLower = strings.ToLower(toContinue)
	if toContinueLower == "yes" || toContinueLower == "y" {
		scoreUser = 0
		scoreComputer = 0
		stonePaperScissor()
	} else if toContinueLower == "no" || toContinueLower == "n" {
		fmt.Printf("\n\t\t\tThank you for playing!\n\n")
		os.Exit(0)
	} else {
		fmt.Println("\t\t\t\nInvalid Input. Say Yes or No!")
		playAgain()
	}
}
