package menu

import (
	"fmt"
	"hangman/game"
	"os"

	"github.com/fatih/color"
)

func Menu() {
	var choix int

	fmt.Print("\033[H\033[2J")
	color.Blue("\\----------Menu Principal----------/")
	color.Blue("     1- Pour lancer une partie")
	color.Blue("     2- Pour voir tout les mots ")
	color.Blue("    \\--------------------------/")
	fmt.Scan(&choix)

	for choix != 1 && choix != 2 && choix != 0 {
		fmt.Println("Veuillez entrer 1 ou 2 !")
		fmt.Scan(&choix)
	}

	switch choix {
	case 1:
		game.TwoPlayers()
		game.Display()
	case 2:
		fmt.Print("\033[H\033[2J")
		data, _ := os.ReadFile("./random/words.txt")
		fmt.Println(string(data))
		fmt.Println("Appuyez sur 0 pour quitter !")
		fmt.Scan(&choix)
		switch choix {
		case 0:
			Menu()
		}
	}
}
