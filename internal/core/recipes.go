package core

import (
	"encoding/json"
	"os"
)

// Recipe represents a crafting recipe, which defines the required input items,
// the resulting output item, and the time it takes to craft.
type Recipe struct {
	// Id is the unique identifier of the recipe.
	Id        int //it is used by internal system, for fast processing
	// Idname is the string identifier of the recipe.
	Idname    string
	// Input is a map of item names to the required count for crafting.
	Input     Items
	// Output is the item produced by the recipe.
	Output    Item
	// CraftTime is the time in ticks required to craft the recipe.
	CraftTime int //in ticks
}

// LoadRecipes loads recipe definitions from a JSON file.
// It panics if the file cannot be read or parsed.
func LoadRecipes(filePath string) []Recipe {
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var recipes []Recipe
	err = json.Unmarshal(jsonFile, &recipes)
	if err != nil {
		panic(err)
	}
	return recipes
}
