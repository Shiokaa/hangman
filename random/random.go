package random

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

func Words() string {
	data, err := os.ReadFile("./random/words.txt") // lire le fichier text.txt
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}

func RandomWord() string {
	n := rand.IntN(313)
	word := ""
	str := strings.Split(Words(), "\n")
	for i, elem := range str {
		if i == n {
			word = elem
		}
	}
	return word
}
