package core

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Add(o Vec2) Vec2 {
	return Vec2{v.X + o.X, v.Y + o.Y}
}

type Task struct {
}

type BaseE struct {
	t Task
}

type Item struct {
	name     string
	id       int
	maxstack int
}

type Recipe struct {
	id     int
	input  Item
	output Item
}

type Production struct {
}

type Inventory struct {
}

type Player struct {
	BaseE
}

type RType struct {
	name string
}

type RPatch struct {
	size  Vec2
	itype RType
	rleft int64
}

type Map struct {
	size Vec2
	rps  []RPatch
}

type Building struct {
	BaseE
}

func (b *BaseE) Update() {

}
