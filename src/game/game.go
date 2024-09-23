package game

import (
	"hangman/random"
	"math/rand/v2"
)

func Test() {
	word := ""
	n := rand.IntN(len(random.RandomWord()) - 1)
	for i, char := range random.RandomWord() {
		if i != n {
			word += "_"
		} else {
			word += string(char)
		}
	}
}
