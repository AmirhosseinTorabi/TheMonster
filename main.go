package main

import "MiladSamani/interaction"

var currentRound = 0

func main() {
	startGame()
	winner := ""
	for winner == "" {
		execute()
	}
	endGame()
}

func startGame() {
	interaction.Print()
}

func execute() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0
	interaction.ShowAvalibaleAction(isSpecialRound)
	return ""
}

func endGame() {

}
