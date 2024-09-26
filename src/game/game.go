package game

import (
	"fmt"
	"hangman/random"
	"math/rand/v2"
	"time"
)

var Word string = random.RandomWord()
var HiddenWord []rune = CreateSlice()
var Counter int = 6

func Display() {
	var choix int
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Bienvenu dans le jeu du pendu ! \n\n")
	time.Sleep(3 * time.Second)
	fmt.Printf("Voici le mot mystère !\n\n")
	fmt.Println(string(HiddenWord))
	time.Sleep(3 * time.Second)
	for Word != string(HiddenWord) {
		fmt.Print("\033[H\033[2J")
		fmt.Println(Word)
		fmt.Println(string(HiddenWord))
		fmt.Printf("\n1- Pour entrer une lettre !")
		fmt.Printf("\n2- Pour entrer un mot !\n\n")
		fmt.Scan(&choix)
		switch choix {
		case 1:
			fmt.Printf("\nChoisisez une lettre ! ")
			FindLetter()
		case 2:
			fmt.Printf("\nEcrivez votre mot ! ")
			FindWord()
		}
	}
}

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
	fmt.Scan(&choix)
	for !(choix >= "a" && choix <= "z" || choix >= "A" && choix <= "Z") {
		fmt.Printf("\nVeuillez entrer une lettre ")
		fmt.Scan(&choix)
	}
	if len(choix) > 1 {
		fmt.Printf("\nVeuillez entrer UNE lettre ")
		fmt.Scan(&choix)
	}
	for iWord, charWord := range Word {
		if choix == string(charWord) {
			for iSlice := range slice {
				if iWord == iSlice {
					slice[iSlice] = charWord
					HiddenWord = slice
					break
				}
			}
		} else {
			Counter--
		}
	}
	return HiddenWord
}

/* Reste a faire si le joueur a marqué toutes les lettres il gagne. Si la lettre est mauvaise -1 à un compteur de 6, si le joueur atteint 0 il perd.
Il faut afficher le compteur et les lettres déjà utilisé, si le joueur tente de reutiliser une lettre déjà utiliser message erreur.
Si il perd évolution de l'ASCII. S'il propose un mot et qu'il a juste messsage victoire, s'il a faux -2 au compteur*/

func FindWord() []rune {
	var choix string
	fmt.Scan(&choix)
	if choix == Word {
		fmt.Printf("Bien joué c'est le bon mot !\n\n")
	} else {
		Counter--
		fmt.Printf("Ce n'était pas le bon mot !\n\n")
	}
	return HiddenWord
}
