package menu

import (
	"fmt"
	"hangman/game"
	"os"
	"time"
)

func Menu() {
	fmt.Print("\033[H\033[2J")
	var choix int
	fmt.Println("\\----------Menu Principal----------/")
	fmt.Println("     1- Pour lancer une partie")
	fmt.Println("     2- Pour voir tout les mots ")
	fmt.Println("    \\--------------------------/")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Bienvenu dans le jeu du pendu !\n\n")
		time.Sleep(2 * time.Second)
		game.PlayerChoseChar()
	case 2:
		data, _ := os.ReadFile("./random/words.txt")
		fmt.Println(string(data))
	}
}
