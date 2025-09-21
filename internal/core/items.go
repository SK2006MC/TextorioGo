package core

import (
	"encoding/json"
	"os"
)

type Item struct {
	name     string
	Id       int32
	maxstack int
}

type Items map[string]int

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
