package core

type Player struct {
	BaseE
	crafting            *Recipe
	craftingTime        int //ticks left
	craftingSpeedFactor float64
	craftingQueue       CraftingQueue
	inv                 Inventory
}

type CraftingQueue struct {
	recipes []Recipe
	ticks   int
}

func (p *Player) Craft(r *Recipe) {
}
