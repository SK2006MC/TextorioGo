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
	var te, sec int64
	lastUpdateTime = time.Now()

	for {
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(lastUpdateTime)
		lastUpdateTime = currentTime
		tickAccumulator += elapsedTime

		// This loop runs the game logic as many times as needed to catch up
		for tickAccumulator >= TickRate {
			game1.Update()
			te++
			//	fmt.Printf("Ticks: %d\r", te)
			tickAccumulator -= TickRate
			if te%60 == 0 {
				sec++
				fmt.Printf("Secs:%d\r", sec)
			}
		}

		// Check for non-blocking user input
		select {
		case input := <-inputChan:
			r := game1.ProcessCommand(input)
			if r == -1 {
				//game1.Save()
				return
			}
		default:
			// No input, continue to the next tick
		}

		time.Sleep(time.Millisecond) // Prevent busy-waiting
	}
}
