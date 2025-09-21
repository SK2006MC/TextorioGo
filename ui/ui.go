package ui

type Dimensions struct {
	Height int
	Width  int
}

type BaseUI interface {
	Draw()
}
