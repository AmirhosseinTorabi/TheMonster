package interaction

import "fmt"

func Print() {
	fmt.Println("===The Monster Game===")
	fmt.Println("Start a New Game...!")
}

func ShowAvalibaleAction(isSpecialAttackIsAvaliable bool)  {
	fmt.Println("Please Choose Your Action")
	fmt.Println("-------------------------")
	fmt.Println("(1)Attack Monster")
	fmt.Println("(2)Heal")
	if isSpecialAttackIsAvaliable {
		fmt.Println("(3)Special Attack")
	}
}