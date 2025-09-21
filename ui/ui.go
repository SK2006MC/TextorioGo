package ui

// Dimensions represents the dimensions (height and width) of a UI element.
type Dimensions struct {
	// Height is the height of the UI element.
	Height int
	// Width is the width of the UI element.
	Width  int
}

// BaseUI is the interface that all UI implementations must satisfy.
type BaseUI interface {
	// Draw redraws the UI.
	Draw()
}
