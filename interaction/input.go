package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoose(isSpecialAttackIsAvaliable bool) string {
	for {
		playerChoice, _ := getPlayerInput()
		if playerChoice == "1" {
			return "ATTACK"
		} else if playerChoice == "2" {
			return "HEAL"
		} else if playerChoice == "3" && isSpecialAttackIsAvaliable {
			return "SPECIAL ATTACK"
		}
		fmt.Println("ERROR! PLEASE TRY AGAIN!")
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your Choice : ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	userInput = strings.Replace(userInput, "\n" ,"", -1)
	return userInput, nil
}
