package ui

// Menu represents the Menu ui of the game
type Menu struct {
	options []Option
}

// Option represents a Option in menu
type Option struct {
	onClick     func()
	name        string
	description string
}

func NewOption(name, description string, onClick func()) *Option {
	return &Option{
		name:        name,
		description: description,
		onClick:     onClick,
	}
}

func (m *Menu) AddOption(op Option) {
	m.options = append(m.options, op)
}
