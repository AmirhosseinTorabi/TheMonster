package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader  = bufio.NewReader(os.Stdin)

func getPlayerChoose(isSpecialAttackIsAvaliable bool) {

}

func getPlayerInput()(string , error) {
	fmt.Print("Your Choice : ")
	userInput , err := reader.ReadString('\n')
	if err != nil {
		return "" , err
	}
	userInput = strings.Replace(userInput , "" , "\n" , -1)
	return userInput , nil
}