package core

import (
	"fmt"
	"strings"
)

type Game struct {
	activeEntities []BaseE
	output1        string
	name           string
	player         Player
	gmap           Map
	tickElapsed    int64
	pr             Production
	lrecipes       []Recipe
	litems         []Item
	lbuildings     []Building
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Write(msg string) {

}

func (g *Game) Update() {
	var uds int
	for i, e := range g.activeEntities {
		e.Update()
		uds += i
	}
}

func (g *Game) Save(filename string) error {
	// Implement save functionality here use binary serialization
	return nil
}

func (g *Game) Load(filename string) error {
	// Implement load functionality here use binary serialization

	return nil
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
