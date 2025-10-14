package core

// Map represents the game map, including its size and resource patches.
type Map struct {
	size Vec2
	rps  []RPatch
}

func NewMap(size Vec2) *Map {

	return &Map{
		size: size,
		rps:  nil,
	}
}
