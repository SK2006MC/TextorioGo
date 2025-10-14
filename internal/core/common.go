package core

// Vec2 represents a 2D vector with X and Y coordinates.
type Vec2 struct {
	// X is the x-coordinate of the vector.
	X float64
	// Y is the y-coordinate of the vector.
	Y float64
}

// Add adds two Vec2 vectors and returns the resulting vector.
func (v Vec2) Add(o Vec2) Vec2 {
	return Vec2{v.X + o.X, v.Y + o.Y}
}

// Task represents a task to be performed by a game entity.
type Task struct {
	run any
}

// BaseE is the base struct for all entities in the game.
type BaseE struct {
	t Task
}

// Production represents the factorio line chart data.
type Production struct {
	ItemsProduced []int64
	ItemsUsed     []int64
	ItemsWasted   []int64
	PowerProduced []int64
}

// RType represents a type of resource.
type RType struct {
	Name string
}

// RPatch represents a resource patch on the game map.
type RPatch struct {
	Size  Vec2
	IType RType
	RLeft int64
}

// Building represents a building in the game.
type Building struct {
	BaseE
}

// Update is the method called on every game tick to update the entity's state.
func (b *BaseE) Update() {

}
