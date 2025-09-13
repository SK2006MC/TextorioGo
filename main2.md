```go
package main

import (
	"Textorio/internal/core"
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2" // Import tcell
)

const (
	TickRate = time.Second / 60
)

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
		SetFieldWidth(30).
		SetDoneFunc(func(key tcell.Key) { // Use tcell.Key
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
	game1 := core.NewGame()

	var lastUpdateTime time.Time
	var tickAccumulator time.Duration
	var te, sec int64
	lastUpdateTime = time.Now()

	for {
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(lastUpdateTime)
		lastUpdateTime = currentTime
		tickAccumulator += elapsedTime

		for tickAccumulator >= TickRate {
			game1.Update()
			te++
			output.Write([]byte(fmt.Sprintf("Tick: %d\r", te)))
			tickAccumulator -= TickRate
			if te%60 == 0 {
				sec++
				output.Write([]byte(fmt.Sprintf("Sec: %d\r", sec)))
			}
		}

		select {
		case input := <-inputChan:
			r := game1.ProcessCommand(input)
			if r == -1 {
				return
			}
		default:
		}

		time.Sleep(time.Millisecond)
	}
}
```
