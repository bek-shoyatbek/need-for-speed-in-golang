package main

import (
	src "Need_for_speed_1.0.0/src"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	src.Start("Game Started")
	time.Sleep(time.Second * 2)
	src.Clean()
	src.Start("Choose a level you want")
	var level int
	_, err := fmt.Scan(&level)
	if err != nil {
		panic(err)
		return
	}

	var move int
	go src.Race(level, &move, &wg)
	go src.GetInput(&move, &wg)

	wg.Wait()

}
