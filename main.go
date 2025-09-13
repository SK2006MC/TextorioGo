package main

import (
	"Textorio/internal/core"
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const TickRate = time.Second / 60

type GameUI struct {
	app         *tview.Application
	outputView  *tview.TextView
	inputField  *tview.InputField
	inputFieldi *tview.Box
}

type GameApp struct {
	gui          *GameUI
	game         *core.Game
	inputChan    chan string
	lastTickTime time.Time
	tickAccum    time.Duration
	tickCount    int64
	secondCount  int64
}

func NewGameApp() *GameApp {
	app := tview.NewApplication()
	game := core.NewGame()
	outputView := tview.NewTextView()
	inputField := tview.NewInputField().
		SetLabel("Enter command: ").
		SetFieldWidth(30)

	inputFieldi := tview.NewBox().
		SetTitle("Input").
		SetBackgroundColor(0xff0000).
		SetBorder(true)

	gui := &GameUI{
		app:         app,
		outputView:  outputView,
		inputField:  inputField,
		inputFieldi: inputFieldi,
	}

	return &GameApp{
		gui:          gui,
		game:         game,
		inputChan:    make(chan string),
		lastTickTime: time.Now(),
	}
}

func (game *GameApp) setupUI() {
	ga := game.gui
	ga.outputView.SetBorder(true).SetTitle("Game Output")
	ga.outputView.SetChangedFunc(func() {
		ga.app.Draw()
	})

	ga.inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			command := ga.inputField.GetText()
			if command != "" {
				result := game.game.ProcessCommand(command)
				if result == -1 {
					ga.app.Stop()
					return
				}
				fmt.Fprintf(ga.outputView, "You entered: %s\n", command)
				ga.inputField.SetText("")
			}
		}
	}).SetBorder(true).SetTitle("Inputi")

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(ga.outputView, 0, 3, false).
		AddItem(ga.inputField, 0, 1, true)

	ga.app.SetRoot(flex, true)
}

func (ga *GameApp) runGameLoop() {
	for {
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(ga.lastTickTime)
		ga.lastTickTime = currentTime
		ga.tickAccum += elapsedTime

		for ga.tickAccum >= TickRate {
			ga.game.Update()
			ga.tickCount++
			ga.tickAccum -= TickRate

			if ga.tickCount%60 == 0 {
				ga.secondCount++
				fmt.Fprintf(ga.gui.outputView, "Tick: %d, Sec: %d\r", ga.tickCount, ga.secondCount)
			}
		}

		time.Sleep(time.Millisecond)
	}
}

func (ga *GameApp) Run() error {
	ga.setupUI()
	go ga.runGameLoop()
	return ga.gui.app.Run()
}

func main() {
	gameApp := NewGameApp()
	if err := gameApp.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
