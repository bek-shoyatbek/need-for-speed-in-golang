package src

import (
	"github.com/eiannone/keyboard"
	"sync"
)

func GetInput(move *int, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if char == 97 {
			if *move == 0 {
				*move = 3
				stopGameOver()
			}
			*move--
		} else if char == 100 {
			if *move == 3 {
				*move = 0
				stopGameOver()
			}
			*move++
		}
	}
}
