package tview

import (
	"Textorio/config"
	"Textorio/internal/core"
	"Textorio/ui"

	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Dimensions is an alias for the ui.Dimensions struct.
type Dimensions = ui.Dimensions

// UIConfig holds the configuration for the UI.
type UIConfig struct {
	// Title is the title of the UI window.
	Title           string
	// Border is a flag indicating whether the UI should have a border.
	Border          bool
	// OutputDim is the dimensions of the output view.
	OutputDim       Dimensions
	// InputDim is the dimensions of the input view.
	InputDim        Dimensions
	// BackgroundColor is the background color of the UI.
	BackgroundColor tcell.Color
}

// GameUI holds the tview UI components for the game.
type GameUI struct {
	app        *tview.Application
	outputView *tview.TextView
	inputField *tview.InputField
	button     *tview.Button
	outputFlex *tview.Flex
}

// GameApp is the main application struct, holding the UI, game state, and timing information.
type GameApp struct {
	gui          *GameUI
	game         *core.Game
	lastTickTime time.Time
	tickAccum    time.Duration
	tickCount    int64
	secondCount  int64
	tickRate     time.Duration
}

// NewGameUI creates and returns a new GameUI instance.
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

// SetupUI configures the UI components based on the provided UIConfig.
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

// ProcessCommand processes a command string entered by the player.
func (gApp *GameApp) ProcessCommand(command string) {
	result := gApp.game.ProcessCommand(command)
	if result == -1 {
		gApp.gui.app.Stop()
		return
	}
	fmt.Fprintf(gApp.gui.outputView, "You entered: %s\n", command)
}

// UpdateOutput updates the game's output view with a new message.
func (gui *GameUI) UpdateOutput(message string) {
	fmt.Fprintf(gui.outputView, "%s\r", message)
	gui.app.Draw()
}

// NewGameApp creates and returns a new GameApp instance.
func NewGameApp(uiConfig UIConfig, tickRate time.Duration) *GameApp {
	if tickRate == 0 {
		tickRate = config.DefaultTickRate
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

// runGameLoop starts the main game loop, which runs on a ticker.
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

// updateOutputWithTickInfo updates the output view with the current tick and second count.
func (ga *GameApp) updateOutputWithTickInfo() {
	message := fmt.Sprintf("Tick: %d, Sec: %d", ga.tickCount, ga.secondCount)
	ga.gui.UpdateOutput(message)
}

// Run starts the application, including the game loop and the tview application.
func (ga *GameApp) Run() error {
	go ga.runGameLoop()
	return ga.gui.app.Run()
}

// CreateDefaultUIConfig creates and returns a default UIConfig.
func CreateDefaultUIConfig() UIConfig {
	return UIConfig{
		Title:           "Input",
		Border:          true,
		BackgroundColor: tcell.NewHexColor(0x000000),
		OutputDim:       Dimensions{Height: 0, Width: 0},
		InputDim:        Dimensions{Height: 1, Width: 30},
	}
}
