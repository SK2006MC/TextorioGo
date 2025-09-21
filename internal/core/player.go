package core

// Player represents the player in the game.
type Player struct {
	BaseE
	crafting            *Recipe
	craftingTime        int //ticks left
	craftingSpeedFactor float64
	craftingQueue       CraftingQueue
	inv                 Inventory
}

// CraftingQueue represents a queue of crafting recipes for the player.
type CraftingQueue struct {
	recipes []Recipe
	ticks   int
}

// Craft adds a recipe to the player's crafting queue.
// Currently, this is a placeholder and does not do anything.
func (p *Player) Craft(r *Recipe) {
}
