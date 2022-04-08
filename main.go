package main

import (
	"github.com/plug-pathomgphong/monster-slayer-golang/interaction"
)

var currentRound = 0

func main() {
	startGame()

	winner := "" // "Player" || " Monster" || ""
	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialAttack := currentRound%3 == 0
	interaction.ShowAvailableActions(isSpecialAttack)          // Show choice
	userChoice := interaction.GetPlayerChoice(isSpecialAttack) // User choose choice from input

	if userChoice == "ATTACK" {

	} else if userChoice == "HEAL" {

	} else {

	}
	return ""
}

func endGame() {

}
