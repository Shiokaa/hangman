package game

import (
	"fmt"
	"hangman/random"
	"math/rand/v2"
	"strings"
	"time"
)

var Word string
var HiddenWord []rune
var Counter int = 6
var LettersAlreadyFound []string
var DiffCounter int
var TwoPlayersCounter int

func Display() {
	var choix int
	var choixDiff int
	if TwoPlayersCounter == 1 {
		text2 := "Bienvenue dans le jeu du pendu ! Trouvez le mot de votre ami !"
		fmt.Print("\033[H\033[2J")

		for _, char := range text2 {
			fmt.Printf("%c", char)
			time.Sleep(40 * time.Millisecond)
		}
	} else {
		text := "Bienvenue dans le jeu du pendu ! Les mots possèdent une majuscule mais pas d'accent, bon jeu à vous !"
		fmt.Print("\033[H\033[2J")

		for _, char := range text {
			fmt.Printf("%c", char)
			time.Sleep(40 * time.Millisecond)
		}
	}
	time.Sleep(3 * time.Second)
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Choisissez votre difficulté !\n\n")
	fmt.Printf("1- Moyen (2 lettres de révélés)\n\n")
	fmt.Printf("2- Difficile (1 lettre de révélé) ")
	fmt.Scan(&choixDiff)
	for choixDiff != 1 && choixDiff != 2 {
		fmt.Printf("Veuillez taper 1 ou 2 ! ")
	}
	if choixDiff == 1 {
		DiffCounter = 2
		HiddenWord = CreateSlice()
	} else {
		DiffCounter = 1
		HiddenWord = CreateSlice()
	}
	for Word != string(HiddenWord) && Counter > 0 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("LE JEU DU PENDU !")
		fmt.Println("\nMot : ", string(HiddenWord))
		fmt.Printf("Essais restants : %d\n", Counter)
		DisplayHangman(Counter)
		letters := strings.Join(LettersAlreadyFound, ", ")

		if len(LettersAlreadyFound) == 1 {
			fmt.Printf("Lettre déjà utilisé : %s", letters)
		}
		if len(LettersAlreadyFound) > 1 {
			fmt.Printf("Lettres déjà utilisés : %s", letters)
		}

		fmt.Printf("\n1- Pour entrer une lettre !")
		fmt.Printf("\n2- Pour entrer un mot !\n\n")

		fmt.Scan(&choix)

		for choix != 1 && choix != 2 {
			fmt.Println("Veuillez entrer 1 ou 2 !")
			fmt.Scan(&choix)
		}

		switch choix {
		case 1:
			fmt.Printf("\nChoisisez une lettre ! ")
			FindLetter()
		case 2:
			fmt.Printf("\nEcrivez votre mot ! ")
			FindWord()
		}
	}

	if Word == string(HiddenWord) {
		fmt.Printf("\nBien joué ! Vous avez trouvé le mot : %s\n", Word)
	} else {
		fmt.Println("\nDésolé, vous avez perdu. Le mot était :", Word)
	}
}

func CreateSlice() []rune {
	var runeSlice []rune
	n := rand.IntN(len(Word) - 1)
	x := rand.IntN(len(Word) - 1)

	for x == n {
		x = rand.IntN(len(Word) - 1)
	}

	if DiffCounter == 1 {
		for i, char := range Word {
			if i == len(Word) {
				break
			}
			if i == n {
				runeSlice = append(runeSlice, char)
			} else {
				runeSlice = append(runeSlice, '_')
			}
		}
	}

	if DiffCounter == 2 {
		for i, char := range Word {
			if i == len(Word) {
				break
			} else if i == n {
				runeSlice = append(runeSlice, char)
			} else if i == x {
				runeSlice = append(runeSlice, char)
			} else {
				runeSlice = append(runeSlice, '_')
			}
		}
	}

	return runeSlice
}

func FindLetter() []rune {
	var choix string
	slice := HiddenWord
	wordFind := false

	fmt.Scan(&choix)
	for !(choix >= "a" && choix <= "z" || choix >= "A" && choix <= "Z") {
		fmt.Printf("\nVeuillez entrer une lettre ")
		fmt.Scan(&choix)
	}
	if len(choix) > 1 {
		fmt.Printf("\nVeuillez entrer UNE lettre ")
		fmt.Scan(&choix)
	}

	for _, letters := range LettersAlreadyFound {
		if choix == letters {
			fmt.Printf("Lettre déjà utilisé ! ")
			fmt.Scan(&choix)
		}
	}

	for iWord, charWord := range Word {
		if choix == string(charWord) {
			wordFind = true
			for iSlice := range slice {
				if iWord == iSlice {
					slice[iSlice] = charWord
					HiddenWord = slice
					break
				}
			}
		}
	}

	if !wordFind {
		LettersAlreadyFound = append(LettersAlreadyFound, choix)
		Counter--

	}

	return HiddenWord
}

func FindWord() []rune {
	var choix string

	fmt.Scan(&choix)
	for !(choix >= "a" && choix <= "z" || choix >= "A" && choix <= "Z") {
		fmt.Printf("\nVeuillez entrer des lettres ")
		fmt.Scan(&choix)
	}

	if choix == Word {
		HiddenWord = []rune(Word)
	} else {
		Counter -= 2
	}

	return HiddenWord
}

func ConvertedWord() []rune {
	randomWord := random.RandomWord()
	var runeSlice []rune

	for _, v := range randomWord {
		if v == 13 {
			break
		} else {
			runeSlice = append(runeSlice, v)
		}
	}

	return runeSlice
}

func TwoPlayers() {
	var choix int
	var wordChoice string
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Selectionner votre mode de jeu !\n\n")
	fmt.Printf("1- 1 joueur\n\n")
	fmt.Printf("2- 2 joueurs\n\n")
	fmt.Scan(&choix)

	for choix != 1 && choix != 2 {
		fmt.Printf("Veuillez taper 1 ou 2 ! ")
	}
	switch choix {
	case 1:
		Word = string(ConvertedWord())
	case 2:
		TwoPlayersCounter = 1
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Joueur 1 : Indiquez le mot que le joueur 2 devra trouver !\n\n")
		fmt.Scan(&wordChoice)

		for !(wordChoice >= "a" && wordChoice <= "z" || wordChoice >= "A" && wordChoice <= "Z") {
			fmt.Printf("\nVeuillez entrer des lettres ")
			fmt.Scan(&choix)
		}

		Word = wordChoice
		fmt.Printf("\nVotre mot a été enregistré ! Bon jeu !")
		time.Sleep(2 * time.Second)
	}
}

func DisplayHangman(counter int) {
	switch counter {
	case 6:
		fmt.Println(`
   +---+
   |   |
       |
       |
       |
       |
=========
		`)
	case 5:
		fmt.Println(`
   +---+
   |   |
   O   |
       |
       |
       |
=========
		`)
	case 4:
		fmt.Println(`
   +---+
   |   |
   O   |
   |   |
       |
       |
=========
		`)
	case 3:
		fmt.Println(`
   +---+
   |   |
   O   |
  /|   |
       |
       |
=========
		`)
	case 2:
		fmt.Println(`
   +---+
   |   |
   O   |
  /|\  |
       |
       |
=========
		`)
	case 1:
		fmt.Println(`
   +---+
   |   |
   O   |
  /|\  |
  /    |
       |
=========
		`)
	case 0:
		fmt.Println(`
   +---+
   |   |
   O   |
  /|\  |
  / \  |
       |
=========
		`)
	}
}
