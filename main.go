package main

import (
	"Textorio/internal/core"
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	TickRate = time.Second / 60
)

func readInput(inputChan chan string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		// This will block until a line is entered,
		// but it's okay because it's in a separate gor
		input, err := reader.ReadString('\n')
		if err == nil {
			inputChan <- input
		}
	}
}

func main() {
	fmt.Println("Starting Textorio...")
	inputChan := make(chan string)
	go readInput(inputChan)
	game1 := core.NewGame()
	//	go game_engine.StartInputListener(game.InputChannel)

	var lastUpdateTime time.Time
	var tickAccumulator time.Duration

	lastUpdateTime = time.Now()

	for {
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(lastUpdateTime)
		lastUpdateTime = currentTime
		tickAccumulator += elapsedTime

		// This loop runs the game logic as many times as needed to catch up
		for tickAccumulator >= TickRate {
			game1.Update()
			tickAccumulator -= TickRate
		}

		// Check for non-blocking user input
		select {
		case input := <-inputChan:
			game1.ProcessInput(input)
		default:
			// No input, continue to the next tick
		}

		time.Sleep(time.Millisecond) // Prevent busy-waiting
	}
}
