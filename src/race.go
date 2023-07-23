package src

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Race(level int, move *int, wg *sync.WaitGroup) {
	defer wg.Done()
	var street [][]string
	fillStreet(&street, level)
	size := len(street)
	for len(street) > 0 {
		time.Sleep(time.Millisecond * 600)
		isGameOver(street[size-1], *move)
		street[size-1][*move] = "@"
		Clean()
		for i := range street {
			fmt.Println(street[i])
		}
		size--
		street = street[:size]
		if size == 0 {
			Start("You won!")
			Restart()
		}
	}
}

func fillStreet(street *[][]string, level int) {
	for i := 0; i < level; i++ {
		way := []string{" ", " ", " ", " "}
		for j := 0; j < rand.Intn(2); j++ {
			randWall := rand.Intn(4)
			way[randWall] = "#"
		}
		*street = append(*street, way)
	}

}
