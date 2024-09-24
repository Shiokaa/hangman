package game

import (
	"fmt"
	"hangman/random"
	"math/rand/v2"
)

var Word string = random.RandomWord()
var HiddenWord []rune = CreateSlice()

func CreateSlice() []rune {
	var runeSlice []rune
	n := rand.IntN(len(Word) - 1)
	for i, char := range Word {
		if i == len(Word)-1 {
			break
		}
		if i == n {
			runeSlice = append(runeSlice, char)
		} else {
			runeSlice = append(runeSlice, '_')
		}
	}
	return runeSlice
}

func FindLetter() []rune {
	var choix string
	slice := HiddenWord
	fmt.Printf("\nEntrez votre lettre !\n")
	fmt.Scan(&choix)
	for iWord, charWord := range Word {
		if choix == string(charWord) {
			for iSlice := range slice {
				if iWord == iSlice {
					slice[iSlice] = charWord
					break
				}
			}
		}
	}
	return slice

	/* Reste a faire si le joueur a marqué toutes les lettres il gagne. Si la lettre est mauvaise -1 à un compteur de 6, si le joueur atteint 0 il perd.
	Il faut afficher le compteur et les lettres déjà utilisé, si le joueur tente de reutiliser une lettre déjà utiliser message erreur.
	Si il perd évolution de l'ASCII. S'il propose un mot et qu'il a juste messsage victoire, s'il a faux -2 au compteur*/
}

func FindWord() {
	var choix string
	fmt.Printf("\nEntrez votre mot !\n")
	fmt.Scan(&choix)
	if choix == Word {
		fmt.Printf("Bien joué c'est le bon mot !\n\n")
	} else {
		fmt.Printf("Ce n'était pas le bon mot !\n\n")
	}
}
