package game

import (
	"fmt"
	"hangman/random"
	"math/rand/v2"
	"strings"
	"time"
)

var Word string = random.RandomWord()

func HiddenWord() string {
	hiddenword := ""
	n := rand.IntN(len(Word) - 1)
	for i, char := range Word {
		if i == len(Word)-1 {
			break
		}
		if i == n {
			hiddenword += string(char)
		} else {
			hiddenword += "_"
		}
	}
	return hiddenword
}

func PlayerChoseChar() {
	hiddenWord := HiddenWord()
	finalResult := ""
	fmt.Println("Voici le mot mystère :\n       ", hiddenWord)
	time.Sleep(2 * time.Second)
	var letter string
	var completeWord string
	var choix int
	fmt.Println("\nQue souhaitez vous faire ?")
	fmt.Println("\n1- Proposer 1 lettre seulement")
	fmt.Printf("2- Proposer 1 mot complet\n\n")
	fmt.Println(Word)
	fmt.Scan(&choix)
	switch choix {
	case 1:
		fmt.Printf("\nEntrez une lettre\n\n")
		fmt.Scan(&letter)
		if !(letter >= "A" && letter <= "Z" || letter >= "a" && letter <= "z") {
			fmt.Println("Saisie invalide réessayez !")
			PlayerChoseChar()
		} else if len(letter) > 1 {
			fmt.Println("Vous ne pouvez proposer qu'une lettre")
			PlayerChoseChar()
		}
		for indexWord, char := range Word {
			hiddenWord = ""
			if letter == string(char) {
				i := strings.Index(Word, string(char))
				for indexHiddenWord, char2 := range hiddenWord {
					str := string(char2)
					if i == indexHiddenWord {
						finalResult = strings.ReplaceAll(hiddenWord, string(str[indexHiddenWord]), letter)
					}
				}
			}
			if len(Word)-1 == indexWord {
				break
			}
		}

	case 2:
		fmt.Println("\nEntrez un mot")
		fmt.Scan(&completeWord)
		if !(completeWord >= "A" && completeWord <= "Z" || completeWord >= "a" && completeWord <= "z") {
			fmt.Println("Saisie invalide réessayez !")
			PlayerChoseChar()
		} else if len(completeWord) > len(Word) {
			fmt.Println("Votre mot est trop grand")
			PlayerChoseChar()
		}
	default:
		fmt.Println("Saisie invalide réessayez !")
		time.Sleep(2 * time.Second)
		PlayerChoseChar()
	}
	fmt.Println(finalResult)
}
