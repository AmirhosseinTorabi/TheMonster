package main

import (
	"MiladSamani/actions"
	"MiladSamani/interaction"
)

var currentRound = 0

func main() {
	startGame()
	winner := ""
	for winner == "" {
		winner = execute()
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
	userChoice := interaction.GetPlayerChoose(isSpecialRound)
	if userChoice == "ATTACK" {
		actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		actions.HealPlayer()
	} else {
		actions.AttackMonster(true)
	}
	actions.AttackPlayer()
	playerHealth , monsterHealth :=actions.GetHealthAmounts()
	if playerHealth<=0 {
		return "Monster"
	} else if monsterHealth <=0 {
		return "Player"
	}
	return ""
}

func endGame() {
}
