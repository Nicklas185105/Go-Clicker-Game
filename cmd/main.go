package main

import (
	"clicker-game/engine"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

func main() {
	game := engine.NewGame()

	// add a click every second in the background
	go func() {
		print("Starting go loop")
		for game.IsRunning() {
			if game.GetClicksPerSecond().GreaterThan(decimal.Zero) {
				game.Tick()
			}

			// wait 1 second
			ticker := time.NewTicker(time.Second)
			<-ticker.C
			ticker.Stop()
		}
	}()

	for game.IsRunning() {
		var input string
		fmt.Scanln(&input)
		game.ProcessInput(input)
	}
}
