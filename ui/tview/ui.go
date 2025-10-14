package tview

import (
	"Textorio/config"
	"Textorio/internal/core"
	"Textorio/ui"
	"os"

	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Dimensions is an alias for the ui.Dimensions struct.
type Dimensions = ui.Dimensions

// UIConfig holds the configuration for the UI.
type UIConfig struct {
	// Title is the title of the UI window.
	Title string
	// Border is a flag indicating whether the UI should have a border.
	Border bool
	// OutputDim is the dimensions of the output view.
	OutputDim Dimensions
	// InputDim is the dimensions of the input view.
	InputDim Dimensions
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
	button.SetBackgroundColor(tcell.ColorGreen).
		SetBorderColor(tcell.ColorDarkRed).
		SetTitleColor(tcell.ColorDarkGreen)

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
		SetPlaceholder("move 0 0").
		SetFieldWidth(uiConfig.InputDim.Width).
		SetFieldBackgroundColor(tcell.ColorDarkGrey)

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

func (gui *GameUI) Print(msg string) {
	fmt.Fprintf(gui.outputView, msg)
}

// ProcessCommand processes a command string entered by the player.
// It returns 0 if the command is empty, -1 if the command is to quit, and 1 otherwise.
func (gApp *GameApp) ProcessCommand(input string) int {
	gui := gApp.gui
	parts := strings.Fields(strings.TrimSpace(input))
	if len(parts) == 0 {
		return 0
	}
	command := parts[0]
	args := parts[1:]

	switch command {
	case "help":
	case "inv":
	case "craft":
		if len(args) == 0 {
			gui.Print("Usage: craft <item_name>")
			return 1
		}
	case "quit":
		gui.Print("Exiting...")
		os.Exit(1)
	default:
		msg := "Unknown command. Type 'help' for a list of commands."
		gui.Print(msg)
	}
	return 1
}

// UpdateOutput updates the game's output view with a new message.
func (gui *GameUI) UpdateOutput(message string) {
	gui.outputView.Clear()
	gui.Print(fmt.Sprintf("%s\n", message))
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

func setUpProcessCmd(game *GameApp) {
	gameUI := game.gui
	gameUI.inputField.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			command := gameUI.inputField.GetText()
			gameUI.inputField.SetText("")
			game.ProcessCommand(command)
		case tcell.KeyESC:
			gameUI.inputField.SetText("")
		}
	})
}

// runGameLoop starts the main game loop, which runs on a ticker.
func (ga *GameApp) runGameLoop() {
	setUpProcessCmd(ga)
	ticker := time.NewTicker(ga.tickRate)
	defer ticker.Stop()
	game := ga.game
	for range ticker.C {
		game.Update()
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
	return ga.gui.app.EnableMouse(true).EnablePaste(true).Run()
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
