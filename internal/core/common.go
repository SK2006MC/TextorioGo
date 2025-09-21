package core

var iid int = 0

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Add(o Vec2) Vec2 {
	return Vec2{v.X + o.X, v.Y + o.Y}
}

type Task struct {
	run interface {
	}
}

type BaseE struct {
	t Task
}

type Production struct {
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
