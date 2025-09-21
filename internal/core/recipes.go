package core

import (
	"encoding/json"
	"os"
)

type Recipe struct {
	Id        int //it is used by internal system, for fast processing
	Idname    string
	Input     Items
	Output    Item
	CraftTime int //in ticks
}

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
