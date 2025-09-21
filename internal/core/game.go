package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"strings"
)

// Game represents the main game state, including all entities, the map, and player information.
type Game struct {
	activeEntities []BaseE
	gmap           Map
	lbuildings     []Building
	litems         []Item
	lrecipes       []Recipe
	name           string
	output1        string
	player         Player
	pr             Production
	tickElapsed    int64
}

// NewGame creates and returns a new Game instance.
func NewGame() *Game {
	return &Game{}
}

// Write sends a message to the game's output.
// Currently, this is a placeholder and does not do anything.
func (g *Game) Write(msg string) {

}

// Update is called on every game tick to update the state of all active entities.
func (g *Game) Update() {
	var uds int
	for i, e := range g.activeEntities {
		e.Update()
		uds += i
	}
}

// Save serializes and saves the current game state to a file.
// It uses gob encoding to write the game state.
func (g *Game) Save(filename string) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(g)
	if err != nil {
		fmt.Println("Error encoding game state:", err)
	}
	err = os.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error writing save file:", err)
		return err
	}
	return nil
}

// Load loads a game state from a file and deserializes it into the Game object.
// It uses gob decoding to read the game state.
func (g *Game) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading save file:", err)
		return err
	}
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err = dec.Decode(g)
	if err != nil {
		fmt.Println("Error decoding game state:", err)
		return err
	}
	return nil
}

// ProcessCommand processes a command string entered by the player.
// It returns 0 if the command is empty, -1 if the command is to quit, and 1 otherwise.
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
