package biblio

import (
	"log"
	"math/rand/v2"
	"os"
)

var alphabet = [...]string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	"q", "w", "e", "r", "t", "y", "u", "i", "o", "p",
	"a", "s", "d", "f", "g", "h", "j", "k", "l", "z",
	"x", "c", "v", "b", "n", "m",
	"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P",
	"A", "S", "D", "F", "G", "H", "J", "K", "L", "Z",
	"X", "C", "V", "B", "N", "M",
	" ", ",", ".", ":", ";", "!", "?",
}

func book(filename string) {
	var text string
	limit := len(alphabet)
	for p := 1; p <= 512; p++ {
		var page string
		for l := 1; l <= 50; l++ {
			var line string
			for c := 1; c <= 80; c++ {
				line += alphabet[rand.IntN(limit)]
			}
			line += "\n"
			page += line
		}
		page += "\n"
		text += page
	}

	err := os.WriteFile(filename, []byte(text), 0666)

	if err != nil {
		log.Fatal(err)
	}
}
