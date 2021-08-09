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
	myFigure := figure.NewFigure("MONSTER", "", true)
	myFigure.Print()
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game...")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("=========================")
	fmt.Println("Please choose your action")
	fmt.Println("=========================")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")
	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	myFigure := figure.NewColorFigure("END GAME", "", "red", true)
	myFigure.Print()
	fmt.Println("-------------------------")
	fmt.Printf("%v WON!\n", winner)
}

func PrintRoundStatictis(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage!\n ", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player Performed the Strong Attack for %v damage!\n ", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player Healed for : %v \n", roundData.PlayerHealth)
	}
	fmt.Printf("Monster Attacked Player for %v damage!\n ", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health : %v\n ", roundData.PlayerHealth)
	fmt.Printf("Monster Health : %v \n ", roundData.MonsterHealth)
}
func WriteToFile(rounds *[]RoundData) {
	//file, err := os.Create("gamelog.txt") //for go run
	pathEX, err := os.Executable()
	if err != nil {
		fmt.Println("Writing a log file failed.Exiting.")
	}
	pathEX = filepath.Dir(pathEX)
	file, err := os.Create(pathEX + "/gamelog.txt")
	if err != nil {
		fmt.Println("Saving a log file failed. Exiting.")
		return
	}
	for index, value := range *rounds {
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
	file.Close()
	fmt.Println("--Wrote Data To Log!--")
}
