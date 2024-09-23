package game

import (
	"fmt"
	"hangman/random"
	"math/rand/v2"
	"strings"
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

func PlayerChoseChar() string {
	newHiddenWord := Hiddenword()
	var choix string
	fmt.Scan(&choix)
	for _, char := range newHiddenWord {
		if choix == string(char) {
			fmt.Println("Vous avez déjà trouvé cette lettre")
			PlayerChoseChar()
		} else {
			for _, char := range Word {
				if strings.Contains(Word, choix) {
					strings.Replace(newHiddenWord, "_", string(char), -1)
					fmt.Println("bravo")
				}
			}
		}
	}
	return newHiddenWord
}
