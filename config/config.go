package config

import "time"

const (
	// DefaultTickRate is the default rate at which the game state is updated.
	// It is set to 60 ticks per second.
	DefaultTickRate = time.Second / 60
)
