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

const DefaultTickRate = time.Second / 60

type UIConfig struct {
	Title           string
	Border          bool
	BackgroundColor tcell.Color
	OutputHeight    int
	InputHeight     int
	InputWidth      int
}

type GameUI struct {
	app        *tview.Application
	outputView *tview.TextView
	inputField *tview.InputField
	inputBox   *tview.Box
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

func NewGameUI(uiConfig UIConfig) *GameUI {
	app := tview.NewApplication()
	outputView := tview.NewTextView()
	inputField := tview.NewInputField().
		SetLabel("Enter command: ").
		SetFieldWidth(uiConfig.InputWidth)

	inputBox := tview.NewBox().
		SetTitle(uiConfig.Title).
		SetBackgroundColor(uiConfig.BackgroundColor).
		SetBorder(uiConfig.Border)

	return &GameUI{
		app:        app,
		outputView: outputView,
		inputField: inputField,
		inputBox:   inputBox,
	}
}

func NewGameApp(uiConfig UIConfig, tickRate time.Duration) *GameApp {
	if tickRate == 0 {
		tickRate = DefaultTickRate
	}

	gameUI := NewGameUI(uiConfig)
	game := core.NewGame()

	return &GameApp{
		gui:          gameUI,
		game:         game,
		inputChan:    make(chan string),
		lastTickTime: time.Now(),
		tickRate:     tickRate,
	}
}

func (ga *GameApp) setupUI(uiConfig UIConfig) {
	gui := ga.gui
	setupOutputView(gui)
	setupInputField(gui, ga, uiConfig)
	setupFlex(gui, uiConfig)
	ga.gui.app.SetRoot(gui.flex, true)
}

func setupOutputView(gui *GameUI) {
	gui.outputView.
		SetBorder(true).
		SetTitle("Game Output")
	gui.outputView.SetChangedFunc(func() {
		gui.app.Draw()
	})
}

func setupInputField(gui *GameUI, ga *GameApp, uiConfig UIConfig) {
	gui.inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			handleInput(gui, ga)
		}
	}).SetBorder(true).SetTitle("Input")
}

func handleInput(gui *GameUI, ga *GameApp) {
	command := gui.inputField.GetText()
	if command != "" {
		result := ga.game.ProcessCommand(command)
		if result == -1 {
			gui.app.Stop()
			return
		}
		fmt.Fprintf(gui.outputView, "You entered: %s\n", command)
		gui.inputField.SetText("")
	}
}

func setupFlex(gui *GameUI, uiConfig UIConfig) {
	gui.flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gui.outputView, uiConfig.OutputHeight, 1, false).
		AddItem(gui.inputField, uiConfig.InputHeight, 1, true)
}

func (ga *GameApp) runGameLoop() {
	for {
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(ga.lastTickTime)
		ga.lastTickTime = currentTime
		ga.tickAccum += elapsedTime

		processTicks(ga)
		time.Sleep(time.Millisecond)
	}
}

func processTicks(ga *GameApp) {
	for ga.tickAccum >= ga.tickRate {
		ga.game.Update()
		ga.tickCount++
		ga.tickAccum -= ga.tickRate

		if ga.tickCount%60 == 0 {
			ga.secondCount++
			updateOutputWithTickInfo(ga)
		}
	}
}

func updateOutputWithTickInfo(ga *GameApp) {
	fmt.Fprintf(ga.gui.outputView, "Tick: %d, Sec: %d\r", ga.tickCount, ga.secondCount)
}

func (ga *GameApp) Run(uiConfig UIConfig) error {
	ga.setupUI(uiConfig)
	go ga.runGameLoop()
	return ga.gui.app.Run()
}

func main() {
	uiConfig := createDefaultUIConfig()
	gameApp := NewGameApp(uiConfig, 0)
	runApplication(gameApp, uiConfig)
}

func createDefaultUIConfig() UIConfig {
	return UIConfig{
		Title:           "Input",
		Border:          true,
		BackgroundColor: tcell.NewHexColor(0xff0000),
		OutputHeight:    3,
		InputHeight:     1,
		InputWidth:      30,
	}
}

func runApplication(gameApp *GameApp, uiConfig UIConfig) {
	if err := gameApp.Run(uiConfig); err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
```
