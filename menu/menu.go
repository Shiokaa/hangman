package menu

import (
	"fmt"
	"os"
)

func Menu() {
	var choix int
	fmt.Println("\\----------Menu Principal----------/")
	fmt.Println("     1- Pour lancer une partie")
	fmt.Println("     2- Pour voir tout les mots ")
	fmt.Println("    \\--------------------------/")
	fmt.Scan(&choix)
	switch choix {
	case 1:
	case 2:
		data, _ := os.ReadFile("./random/words.txt")
		fmt.Println(string(data))
	}
}
