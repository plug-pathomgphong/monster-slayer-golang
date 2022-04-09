package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreeting() {
	asciiFigure := figure.NewFigure("Monster slayer", "", true)
	asciiFigure.Print()
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("-------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func (roundData *RoundData) PrintStatistics() {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack against monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v HP.\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health: %v \n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v \n", roundData.MonsterHealth)
}

func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	fmt.Println("GAME OVER")
	asciiFigure := figure.NewColorFigure("GAME OVER", "", "red", true)
	asciiFigure.Print()
	fmt.Println("-------------------------")
	if winner != "Draw" {
		fmt.Printf("%v won!\n", winner)
	} else {
		fmt.Printf("%v!!!\n", winner)
	}

}

func WriteLogFile(roundData *[]RoundData, winner string) {
	exPath, err := os.Executable()

	if err != nil {
		fmt.Println("Writing into log file failed. Exiting.")
		return
	}

	exPath = filepath.Dir(exPath)

	file, err := os.Create(exPath + "/gamelog.txt") // For "go build"
	// file, err := os.Create("gamelog.txt") // For "go run"

	if err != nil {
		fmt.Println("Saving a log file failed. Exiting.")
		return
	}

	for index, value := range *roundData {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value":     fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player Health":         fmt.Sprint(value.PlayerHealth),
			"Monster Health":        fmt.Sprint(value.MonsterHealth),
		}

		logLine := fmt.Sprintln(logEntry)
		_, err = file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writing into log file failed. Exiting.")
			continue
		}
	}
	var result string
	if winner == "Draw" {
		result = fmt.Sprintln("Result: Draw")
	} else {
		result = fmt.Sprintf("Result: %v win!!!", winner)
	}
	_, err = file.WriteString(result)
	if err != nil {
		fmt.Println("Writing into log file failed. Exiting.")
	}
	file.Close()
	fmt.Println("Wrote data to log!")
}
