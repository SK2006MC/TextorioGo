package main

import (
	"Textorio/internal/core" // Assuming this is your game logic package
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const TickRate = time.Second / 60

// readInput continuously reads from stdin and sends non-empty lines to the inputChan.
// This function is currently commented out in the main game loop but can be useful
// for background input processing if needed.
func readInput(inputChan chan string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err == nil {
			// Trim whitespace and newline characters before sending
			trimmedInput := input[:len(input)-1] // Remove the newline character
			if trimmedInput != "" {
				inputChan <- trimmedInput
			}
		}
	}
}

// setupUI initializes and configures the tview UI elements.
func setupUI(app *tview.Application, game *core.Game, output *tview.TextView, inputChan chan string) *tview.InputField {
	// Configure the TextView for displaying game output
	output.SetBorder(true).SetTitle("Game Output")
	output.SetChangedFunc(func() {
		app.Draw() // Redraw the application when the TextView content changes
	})

	// Configure the InputField for user commands
	inputField := tview.NewInputField().
		SetLabel("Enter command: ").
		SetFieldWidth(30)

	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			command := inputField.GetText()
			if command != "" {
				// Process the command through the game logic
				result := game.ProcessCommand(command)
				if result == -1 { // Assuming -1 is a signal to exit
					app.Stop()
					return
				}
				// Display the command entered in the output TextView
				fmt.Fprintf(output, "You entered: %s\n", command)
				inputField.SetText("") // Clear the input field after processing
			}
		}
	})

	return inputField
}

// runGameLoop manages the game's update logic and tick rate.
func runGameLoop(app *tview.Application, game *core.Game, output *tview.TextView) {
	// Uncomment the following lines if you want to use the background input reader
	// inputChan := make(chan string)
	// go readInput(inputChan)

	var lastUpdateTime time.Time
	var tickAccumulator time.Duration
	var tickCount, secondCount int64
	lastUpdateTime = time.Now()

	for {
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(lastUpdateTime)
		lastUpdateTime = currentTime
		tickAccumulator += elapsedTime

		// Process ticks at a fixed rate
		for tickAccumulator >= TickRate {
			game.Update()
			tickCount++
			tickAccumulator -= TickRate

			// Update display periodically (e.g., every second)
			if tickCount%60 == 0 { // Assuming 60 ticks per second for demonstration
				secondCount++
				fmt.Fprintf(output, "Tick: %d, Sec: %d\r", tickCount, secondCount) // Use \r for overwriting on the same line
			}
		}

		// Handle input from the UI (which is now handled by SetDoneFunc)
		// If using the background input reader, you would use a select statement here:
		// select {
		// case input := <-inputChan:
		// 	result := game.ProcessCommand(input)
		// 	if result == -1 {
		// 		app.Stop()
		// 		return
		// 	}
		// 	fmt.Fprintf(output, "System input: %s\n", input)
		// default:
		// }

		// Small sleep to prevent busy-waiting and allow other goroutines to run
		time.Sleep(time.Millisecond)
	}
}

func main() {
	app := tview.NewApplication()

	// Initialize game and output TextView
	game := core.NewGame()
	outputTextView := tview.NewTextView()

	// Create input channel (currently not used directly by the game loop, but good practice)
	inputChan := make(chan string)

	// Setup UI elements
	inputField := setupUI(app, game, outputTextView, inputChan)

	// Create a flex layout to arrange the TextView and InputField
	// The order in AddItem determines their vertical stacking.
	// AddItem(item, height, width_percentage, focusable)
	flex := tview.NewFlex().
		AddItem(outputTextView, 0, 1, false). // TextView takes available space, not focusable initially
		AddItem(inputField, 1, 0, true)       // InputField takes 1 row, not taking width percentage, focusable

	// Set the root of the application to the flex layout
	app.SetRoot(flex, true)

	// Start the game loop in a separate goroutine
	go runGameLoop(app, game, outputTextView)

	// Run the tview application
	if err := app.Run(); err != nil {
		panic(err)
	}
}
