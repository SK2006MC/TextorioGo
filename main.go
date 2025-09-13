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

type FlexDirection int

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

func NewGameApp(uiConfig UIConfig, tickRate time.Duration) *GameApp {
	app := tview.NewApplication()
	game := core.NewGame()
	outputView := tview.NewTextView()
	inputField := tview.NewInputField().
		SetLabel("Enter command: ").
		SetFieldWidth(uiConfig.InputWidth)

	inputBox := tview.NewBox().
		SetTitle(uiConfig.Title).
		SetBackgroundColor(uiConfig.BackgroundColor).
		SetBorder(uiConfig.Border)

	gui := &GameUI{
		app:        app,
		outputView: outputView,
		inputField: inputField,
		inputBox:   inputBox,
	}

	if tickRate == 0 {
		tickRate = DefaultTickRate
	}

	return &GameApp{
		gui:          gui,
		game:         game,
		inputChan:    make(chan string),
		lastTickTime: time.Now(),
		tickRate:     tickRate,
	}
}

func (ga *GameApp) setupUI(uiConfig UIConfig) {
	gui := ga.gui
	gui.outputView.SetBorder(true).SetTitle("Game Output")
	gui.outputView.SetChangedFunc(func() {
		gui.app.Draw()
	})

	gui.inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
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
	}).SetBorder(true).SetTitle("Input")

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gui.outputView, 0, uiConfig.OutputHeight, false).
		AddItem(gui.inputField, 0, uiConfig.InputHeight, true)

	gui.app.SetRoot(flex, true)
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
				fmt.Fprintf(ga.gui.outputView, "Tick: %d, Sec: %d\r", ga.tickCount, ga.secondCount)
			}
		}

		time.Sleep(time.Millisecond)
	}
}

func (ga *GameApp) Run(uiConfig UIConfig) error {
	ga.setupUI(uiConfig)
	go ga.runGameLoop()
	return ga.gui.app.Run()
}

func main() {
	uiConfig := UIConfig{
		Title:           "Input",
		Border:          true,
		BackgroundColor: tcell.NewHexColor(0xff0000),
		OutputHeight:    3,
		InputHeight:     1,
		InputWidth:      30,
	}

	gameApp := NewGameApp(uiConfig, 0) // Use default tick rate
	if err := gameApp.Run(uiConfig); err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
