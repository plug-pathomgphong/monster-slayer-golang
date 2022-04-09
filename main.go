package main

import (
	"github.com/plug-pathomgphong/monster-slayer-golang/actions"
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

	var playerHealth int
	var monsterHealth int

	if userChoice == "ATTACK" {
		actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		actions.HealPlayer()
	} else {
		actions.AttackMonster(true)
	}

	actions.AttackPlayer()
	playerHealth, monsterHealth = actions.GetHealthAmounts()

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame() {

}
