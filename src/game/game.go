package game

import (
	"fmt"
	"hangman/random"
	"math/rand/v2"
	"strings"
	"time"

	"github.com/fatih/color"
)

var Word string
var HiddenWord []rune
var Counter int = 6
var LettersAlreadyFound []string
var WordAlreadyFound []string
var DiffCounter int
var TwoPlayersCounter int

func Display() {
	var choixDiff int
	if TwoPlayersCounter == 1 {
		text2 := color.BlueString("Bienvenue dans le jeu du pendu ! Trouvez le mot de votre ami !")
		fmt.Print("\033[H\033[2J")

		for _, char := range text2 {
			fmt.Printf("%c", char)
			time.Sleep(40 * time.Millisecond)
		}
	} else {
		text := color.BlueString("Bienvenue dans le jeu du pendu ! Les mots possèdent une majuscule mais pas d'accent, bon jeu à vous !")
		fmt.Print("\033[H\033[2J")

		for _, char := range text {
			fmt.Printf("%c", char)
			time.Sleep(40 * time.Millisecond)
		}
	}
	time.Sleep(3 * time.Second)
	fmt.Print("\033[H\033[2J")
	color.Blue("Choisissez votre difficulté !\n\n")
	color.Green("1- Moyen (2 lettres de révélés)\n\n")
	color.Red("2- Difficile (1 lettre de révélé) ")
	fmt.Scan(&choixDiff)
	for choixDiff != 1 && choixDiff != 2 {
		color.Blue("Veuillez taper 1 ou 2 ! ")
	}
	if choixDiff == 1 {
		DiffCounter = 2
		HiddenWord = CreateSlice()
	} else {
		DiffCounter = 1
		HiddenWord = CreateSlice()
	}
	for Word != string(HiddenWord) && Counter >= 0 {
		fmt.Print("\033[H\033[2J")
		color.Yellow("LE JEU DU PENDU !")
		fmt.Println("\nMot : ", string(HiddenWord))
		fmt.Printf("Essais restants : %d\n", Counter)
		DisplayHangman(Counter)

		if Counter == 0 {
			break
		}

		if len(LettersAlreadyFound) == 1 {
			fmt.Printf("Lettre déjà utilisé : %s", strings.Join(LettersAlreadyFound, ", "))
		}
		if len(LettersAlreadyFound) > 1 {
			fmt.Printf("Lettres déjà utilisées : %s", strings.Join(LettersAlreadyFound, ", "))
		}

		if len(WordAlreadyFound) == 1 {
			fmt.Printf("\nMot déjà utilisé : %s", strings.Join(WordAlreadyFound, ", "))
		}
		if len(WordAlreadyFound) > 1 {
			fmt.Printf("\nMots déjà utilisés : %s", strings.Join(WordAlreadyFound, ", "))
		}

		fmt.Printf("\n1- Pour entrer une lettre ou un mot !\n\n")
		FindLetterOrWord()
	}
	if Word == string(HiddenWord) {
		fmt.Printf("\n\033[32mBien joué ! Vous avez trouvé le mot : %s\033[0m\n", Word)
	} else {
		fmt.Println("\n\033[31mDésolé, vous avez perdu. Le mot était :\033[0m", Word)
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

func FindLetterOrWord() []rune {
	var choix string
	slice := HiddenWord
	letterFind := false

	fmt.Scan(&choix)
	for !(choix >= "a" && choix <= "z" || choix >= "A" && choix <= "Z") {
		fmt.Printf("\nVeuillez entrer que des lettres ")
		fmt.Scan(&choix)
	}

	if len(choix) > 1 {
		if choix == Word {
			HiddenWord = []rune(Word)
		} else {
			WordAlreadyFound = append(WordAlreadyFound, choix)
			Counter -= 2
		}

	} else {
		for _, letters := range LettersAlreadyFound {
			if choix == letters {
				fmt.Printf("Lettre déjà utilisé ! ")
				fmt.Scan(&choix)
			}
		}

		for iWord, charWord := range Word {
			if choix == string(charWord) {
				letterFind = true
				for iSlice := range slice {
					if iWord == iSlice {
						slice[iSlice] = charWord
						HiddenWord = slice
						break
					}
				}
			}
		}

		if !letterFind {
			LettersAlreadyFound = append(LettersAlreadyFound, choix)
			Counter--

		}
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
	color.Blue("Selectionner votre mode de jeu !\n\n")
	color.Blue("1- 1 joueur\n\n")
	color.Blue("2- 2 joueurs\n\n")
	fmt.Scan(&choix)

	for choix != 1 && choix != 2 {
		color.Blue("Veuillez taper 1 ou 2 ! ")
	}
	switch choix {
	case 1:
		Word = string(ConvertedWord())
	case 2:
		TwoPlayersCounter = 1
		fmt.Print("\033[H\033[2J")
		color.Blue("Joueur 1 : Indiquez le mot que le joueur 2 devra trouver !\n\n")
		fmt.Scan(&wordChoice)

		for !(wordChoice >= "a" && wordChoice <= "z" || wordChoice >= "A" && wordChoice <= "Z") {
			color.Blue("\nVeuillez entrer des lettres ")
			fmt.Scan(&choix)
		}

		Word = wordChoice
		color.Blue("\nVotre mot a été enregistré ! Bon jeu !")
		time.Sleep(2 * time.Second)
	}
}

func DisplayHangman(counter int) {
	switch counter {
	case 6:
		color.Red(`
   +---+
   |   |
       |
       |
       |
       |
=========
		`)
	case 5:
		color.Red(`
   +---+
   |   |
   O   |
       |
       |
       |
=========
		`)
	case 4:
		color.Red(`
   +---+
   |   |
   O   |
   |   |
       |
       |
=========
		`)
	case 3:
		color.Red(`
   +---+
   |   |
   O   |
  /|   |
       |
       |
=========
		`)
	case 2:
		color.Red(`
   +---+
   |   |
   O   |
  /|\  |
       |
       |
=========
		`)
	case 1:
		color.Red(`
   +---+
   |   |
   O   |
  /|\  |
  /    |
       |
=========
		`)
	case 0:
		color.Red(`
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
