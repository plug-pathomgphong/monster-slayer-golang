package main

import (
	"github.com/plug-pathomgphong/monster-slayer-golang/actions"
	"github.com/plug-pathomgphong/monster-slayer-golang/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	winner := "" // "Player" || " Monster" || ""
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
	isSpecialAttack := currentRound%3 == 0
	interaction.ShowAvailableActions(isSpecialAttack)          // Show choice
	userChoice := interaction.GetPlayerChoice(isSpecialAttack) // User choose choice from input

	var playerHealth int
	var monsterHealth int
	var playerAttackDmg int
	var playerHealValue int
	var monsterAttacDmg int

	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttacDmg = actions.AttackPlayer()
	playerHealth, monsterHealth = actions.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action:           userChoice,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue:  playerHealValue,
		MonsterAttackDmg: monsterAttacDmg,
	}

	gameRounds = append(gameRounds, roundData)

	// interaction.PrintStatistics(&roundData)
	roundData.PrintStatistics()

	if playerHealth <= 0 && monsterHealth > 0 {
		return "Monster"
	} else if monsterHealth <= 0 && playerHealth > 0 {
		return "Player"
	} else if playerHealth <= 0 && monsterHealth <= 0 {
		return "Draw"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds, winner)
}
