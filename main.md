```go
package main

import (
	"Textorio/internal/core"
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	defaultTickRate = time.Second / 60
	fpsUpdateInterval = 1 * time.Second
)

type Dimensions struct {
	Height, Width int
}

type UIConfig struct {
	Title           string
	Border          bool
	BackgroundColor tcell.Color
	OutputDim       Dimensions
	InputDim        Dimensions
}

type GameUI struct {
	app        *tview.Application
	outputView *tview.TextView
	inputField *tview.InputField
	flex       *tview.Flex
}

type GameApp struct {
	gui          *GameUI
	game         *core.Game
	inputChan    chan string
	lastTickTime time.Time
	tickAccum    time.Duration
	tickCount    int64
	secondCount  int64
	tickRate     time.Duration
}

func NewGameUI(config UIConfig) *GameUI {
	app := tview.NewApplication()
	outputView := tview.NewTextView()
	inputField := tview.NewInputField().
		SetLabel("Enter command: ").
		SetFieldWidth(config.InputDim.Width)

	return &GameUI{
		app:        app,
		outputView: outputView,
		inputField: inputField,
	}
}

func NewGameApp(uiConfig UIConfig, tickRate time.Duration) *GameApp {
	if tickRate == 0 {
		tickRate = defaultTickRate
	}

	gameUI := NewGameUI(uiConfig)
	game := core.NewGame()

	gameUI.outputView.
		SetBorder(true).
		SetTitle("Game Output").
		SetBackgroundColor(uiConfig.BackgroundColor)
	gameUI.outputView.SetChangedFunc(func() {
		gameUI.app.Draw()
	})

	gameUI.inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			command := gameUI.inputField.GetText()
			if command != "" {
				result := game.ProcessCommand(command)
				if result == -1 {
					gameUI.app.Stop()
					return
				}
				fmt.Fprintf(gameUI.outputView, "You entered: %s\n", command)
				gameUI.inputField.SetText("")
			}
		}
	}).SetBorder(true).SetTitle("Input")

	gameUI.flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gameUI.outputView, uiConfig.OutputDim.Height, 1, false).
		AddItem(gameUI.inputField, uiConfig.InputDim.Height, 1, true)

	gameUI.app.SetRoot(gameUI.flex, true)

	return &GameApp{
		gui:          gameUI,
		game:         game,
		inputChan:    make(chan string),
		lastTickTime: time.Now(),
		tickRate:     tickRate,
	}
}

func (ga *GameApp) runGameLoop() {
	ticker := time.NewTicker(ga.tickRate)
	defer ticker.Stop()

	for range ticker.C {
		ga.game.Update()
		ga.tickCount++

		if ga.tickCount%uint64(fpsUpdateInterval.Seconds()) == 0 {
			ga.secondCount++
			ga.updateOutputWithTickInfo()
		}
	}
}

func (ga *GameApp) updateOutputWithTickInfo() {
	fmt.Fprintf(ga.gui.outputView, "Tick: %d, Sec: %d\r", ga.tickCount, ga.secondCount)
	ga.gui.app.Draw()
}

func (ga *GameApp) Run() error {
	go ga.runGameLoop()
	return ga.gui.app.Run()
}

func main() {
	uiConfig := createDefaultUIConfig()
	gameApp := NewGameApp(uiConfig, 0)
	if err := gameApp.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}

func createDefaultUIConfig() UIConfig {
	return UIConfig{
		Title:           "Input",
		Border:          true,
		BackgroundColor: tcell.NewHexColor(0x000000),
		OutputDim:       Dimensions{Height: 0, Width: 0},
		InputDim:        Dimensions{Height: 1, Width: 30},
	}
}
```
