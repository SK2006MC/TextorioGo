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

type Dim struct {
	H, W int
}

type UIConfig struct {
	Title           string
	Border          bool
	BackgroundColor tcell.Color
	OutputDim       Dim
	InputDim        Dim
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
		SetFieldWidth(uiConfig.InputDim.W)

	inputBox := tview.NewBox().
		SetTitle(uiConfig.Title).
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
	gui.outputView.
		SetBorder(true).
		SetTitle("Game Output")
	gui.outputView.SetChangedFunc(func() {
		gui.app.Draw()
	})
	gui.inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			handleInput(gui, ga)
		}
	}).SetBorder(true).SetTitle("Input")
	gui.flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gui.outputView, uiConfig.OutputDim.H, 1, false).
		AddItem(gui.inputField, uiConfig.InputDim.H, 1, true)
	ga.gui.app.SetRoot(gui.flex, true)
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

func (ga *GameApp) runGameLoop() {
	for {
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(ga.lastTickTime)
		ga.lastTickTime = currentTime
		ga.tickAccum += elapsedTime

		for ga.tickAccum >= ga.tickRate {
			ga.game.Update()
			ga.tickCount++
			ga.tickAccum -= ga.tickRate
			if ga.tickCount%60 == 0 {
				ga.secondCount++
				updateOutputWithTickInfo(ga)
			}
		}
		time.Sleep(time.Millisecond)
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
		OutputDim:       Dim{1, 1},
		InputDim:        Dim{1, 30},
	}
}

func runApplication(gameApp *GameApp, uiConfig UIConfig) {
	if err := gameApp.Run(uiConfig); err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
