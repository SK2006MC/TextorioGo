package core

import (
	"encoding/json"
	"os"
)

// Item represents a type of item in the game.
type Item struct {
	name     string
	// Id is the unique identifier of the item.
	Id       int32
	maxstack int
}

// Items is a map that holds the count of each item type, using the item name as the key.
type Items map[string]int

// LoadItems loads item definitions from a JSON file.
// It panics if the file cannot be read or parsed.
func LoadItems(filePath string) Items {
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var items Items
	err = json.Unmarshal(jsonFile, &items)
	if err != nil {
		panic(err)
	}
	return items
}
