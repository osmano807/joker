package plugins

import . "github.com/osmano807/joker/interfaces"

type Plugin interface {
	Name() string
	Init()
	Handle(*InputLine) *OutputLine
}

