```go
package main

import (
	"Textorio/internal/core"
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const TickRate = time.Second / 60

func readInput(inputChan chan string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err == nil {
			inputChan <- input
		}
	}
}

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().SetChangedFunc(func() {
		app.Draw()
	})
	textView.SetBorder(true).SetTitle("Game Output")

	inputField := tview.NewInputField().
		SetLabel("Enter command: ").
		SetFieldWidth(30)

	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			command := inputField.GetText()
			if command != "" {
				textView.Write([]byte(fmt.Sprintf("You entered: %s\n", command)))
				inputField.SetText("")
			}
		}
	})

	flex := tview.NewFlex().AddItem(textView, 0, 1, false).AddItem(inputField, 1, 1, true)

	go runGameLoop(textView)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

func runGameLoop(output *tview.TextView) {
	inputChan := make(chan string)
	go readInput(inputChan)
	game := core.NewGame()

	var lastUpdateTime time.Time
	var tickAccumulator time.Duration
	var tickCount, secondCount int64
	lastUpdateTime = time.Now()

	for {
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(lastUpdateTime)
		lastUpdateTime = currentTime
		tickAccumulator += elapsedTime

		for tickAccumulator >= TickRate {
			game.Update()
			tickCount++
			output.Write([]byte(fmt.Sprintf("Tick: %d\r", tickCount)))
			tickAccumulator -= TickRate

			if tickCount%60 == 0 {
				secondCount++
				output.Write([]byte(fmt.Sprintf("Sec: %d\r", secondCount)))
			}
		}

		select {
		case input := <-inputChan:
			result := game.ProcessCommand(input)
			if result == -1 {
				return
			}
		default:
		}

		time.Sleep(time.Millisecond)
	}
}
```
