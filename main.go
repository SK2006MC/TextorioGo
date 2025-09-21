package main

import (
	"Textorio/ui/tview"
	"fmt"
	"os"
)

// main is the entry point of the application.
// It creates a default UI configuration, initializes a new game application, and runs it.
func main() {
	uiConfig := tview.CreateDefaultUIConfig()
	gameApp := tview.NewGameApp(uiConfig, 0)
	if err := gameApp.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
