package game

import (
	"fmt"
	"hangman/random"
	"math/rand/v2"
)

var Word string = random.RandomWord()

func Hiddenword() string {
	hiddenword := ""
	n := rand.IntN(len(Word) - 1)
	for i, char := range Word {
		if i == n {
			hiddenword += string(char)
		} else {
			hiddenword += "_"
		}
		if i == len(Word)-1 {
			break
		}
	}
	return hiddenword
}

func PlayerChoseChar() {
	var choix string
	fmt.Scan(&choix)
	for _, char := range Word {
		if 
	}
}
