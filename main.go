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
)

type Dimensions struct {
	Height int
	Width  int
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
	button     *tview.Button
	outputFlex *tview.Flex
}

type GameApp struct {
	gui          *GameUI
	game         *core.Game
	lastTickTime time.Time
	tickAccum    time.Duration
	tickCount    int64
	secondCount  int64
	tickRate     time.Duration
}

func NewGameUI() *GameUI {
	app := tview.NewApplication()
	outputView := tview.NewTextView()
	inputField := tview.NewInputField()
	button := tview.NewButton("Process Command")

	return &GameUI{
		app:        app,
		outputView: outputView,
		inputField: inputField,
		button:     button,
	}
}

func (gui *GameUI) SetupUI(uiConfig UIConfig) {
	gui.outputView.
		SetBorder(true).
		SetTitle("Game Output").
		SetBackgroundColor(uiConfig.BackgroundColor)
	gui.outputView.SetChangedFunc(func() {
		gui.app.Draw()
	})

	gui.inputField.
		SetLabel("Enter command: ").
		SetFieldWidth(uiConfig.InputDim.Width)

	inputFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(gui.inputField, 0, 2, false).
		AddItem(gui.button, 0, 1, false)

	gui.outputFlex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gui.outputView, uiConfig.OutputDim.Height, 1, false).
		AddItem(inputFlex, uiConfig.InputDim.Height, 1, false)

	gui.app.SetRoot(gui.outputFlex, true)
}

func (gApp *GameApp) ProcessCommand(command string) {
	result := gApp.game.ProcessCommand(command)
	if result == -1 {
		gApp.gui.app.Stop()
		return
	}
	fmt.Fprintf(gApp.gui.outputView, "You entered: %s\n", command)
}

func (gui *GameUI) UpdateOutput(message string) {
	fmt.Fprintf(gui.outputView, "%s\r", message)
	gui.app.Draw()
}

func NewGameApp(uiConfig UIConfig, tickRate time.Duration) *GameApp {
	if tickRate == 0 {
		tickRate = defaultTickRate
	}

	gameUI := NewGameUI()
	game := core.NewGame()

	gameUI.SetupUI(uiConfig)

	return &GameApp{
		gui:          gameUI,
		game:         game,
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

		if ga.tickCount%60 == 0 {
			ga.secondCount++
			ga.updateOutputWithTickInfo()
		}
	}
}

func (ga *GameApp) updateOutputWithTickInfo() {
	message := fmt.Sprintf("Tick: %d, Sec: %d", ga.tickCount, ga.secondCount)
	ga.gui.UpdateOutput(message)
}

func (ga *GameApp) Run() error {
	go ga.runGameLoop()
	return ga.gui.app.Run()
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

func main() {
	uiConfig := createDefaultUIConfig()
	gameApp := NewGameApp(uiConfig, 0)
	if err := gameApp.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
