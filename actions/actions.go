package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)

var currentHealthPlayer = PLAYER_HEALTH
var currentHealthMONSTER = MONSTER_HEALTH


func AttackMonster(isSpecialAttack bool) {
	minAttack := PLAYER_ATTACK_MIN_DMG
	maxAttack := PLAYER_ATTACK_MAX_DMG
	if isSpecialAttack {
		minAttack = PLAYER_ATTACK_MIN_DMG_SPECIAL
		maxAttack = PLAYER_ATTACK_MAX_DMG_SPECIAL
	}
	dmgHeal := GenerateRandBetween(minAttack, maxAttack)
	currentHealthMONSTER -= dmgHeal
}

func HealPlayer() {
	healValue := GenerateRandBetween(PLAYER_HEAL_MAX, PLAYER_HEAL_MIN)
	healValueDiff := PLAYER_HEALTH - currentHealthPlayer
	if healValueDiff >= healValue {
		currentHealthPlayer += healValue
	} else {
		currentHealthPlayer = PLAYER_HEALTH
	}
}

func AttackPlayer() {
	minAttack := MONSTER_ATTACK_MIN
	maxAttack := MONSTER_ATTACK_MAX
	dmgHeal := GenerateRandBetween(minAttack, maxAttack)
	currentHealthPlayer -= dmgHeal
}

func GetHealthAmounts() (int , int)  {
	return currentHealthPlayer , currentHealthMONSTER
}

func GenerateRandBetween(max, min int) int {
	return randGenerator.Intn(max-min) + min
}
 