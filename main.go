package main

import (
	"Textorio/ui/tview"
	"fmt"
	"os"
)

func main() {
	uiConfig := tview.CreateDefaultUIConfig()
	gameApp := tview.NewGameApp(uiConfig, 0)
	if err := gameApp.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
