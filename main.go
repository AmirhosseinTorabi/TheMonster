package main

import (
	"MiladSamani/actions"
	"MiladSamani/interaction"
)

var currentRound = 0
var roundGame = []interaction.RoundData{}

func main() {
	startGame()
	winner := "" // "Player" || "Monster" || ""
	for winner == "" {
		winner = executeRound()
	}
	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0
	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)
	var playerAttackDmg int
	var playerHealthValue int
	var monsterAttackDmg int
	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealthValue = actions.HealPlayer()
	} else {
		playerAttackDmg = actions.AttackMonster(true)
	}
	monsterAttackDmg = actions.AttackPlayer()
	playerHealth, monsterHealth := actions.GetHealthAmounts()
	roundData := interaction.RoundData{
		Action:        userChoice,
		PlayerHealth:  playerHealth,
		MonsterHealth: monsterHealth,
		PlayerAttackDmg: playerAttackDmg,
		PlayerHealValue: playerHealthValue,
		MonsterAttackDmg: monsterAttackDmg,
	}
	interaction.PrintRoundStatictis(&roundData)
	roundGame =append(roundGame,roundData)
	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}
	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteToFile(&roundGame)
}
