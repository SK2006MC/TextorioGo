package core

import (
	"fmt"
	"strings"
)

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() {
	var uds int
	for i, e := range g.activeEntities {
		e.Update()
		uds += i
	}
}

func (g *Game) ProcessCommand(input string) int {
	parts := strings.Fields(strings.TrimSpace(input))
	if len(parts) == 0 {
		return 0
	}
	command := parts[0]
	args := parts[1:]

	switch command {
	case "help":
		fmt.Println("")
	case "inv":
		fmt.Println("")
	case "craft":
		if len(args) == 0 {
			fmt.Println("Usage: craft <item_name>")
			return -1
		}
	case "quit":
		fmt.Println("Exiting...")
		return -1
	default:
		fmt.Println("Unknown command. Type 'help' for a list of commands.")
	}
	return 1
}
