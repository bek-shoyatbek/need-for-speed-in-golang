package src

import (
	"fmt"
	"os"
)

func stopGameOver() {
	Clean()
	fmt.Println("You lose!")
	os.Exit(1)
}

func isGameOver(street []string, movement int) {
	fmt.Println(street[movement])
	if street[movement] == "#" {
		stopGameOver()
	}
}
