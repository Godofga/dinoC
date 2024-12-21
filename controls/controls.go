package controls

import tea "github.com/charmbracelet/bubbletea"

type direction int

const (
	UP direction = iota
	DOWN
	LEFT
	RIGHT
)

type GameControl interface {
	MoveTo(d direction)
	Back()
	ActionA()
	ActionB()
}

type Game interface {
	tea.Model
}
