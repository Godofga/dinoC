package styles

import "github.com/charmbracelet/lipgloss"

var (
	UnrevealedField = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#3D3634"))
	RevealedField = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4"))
	FlaggedField = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#FFFF00"))
	BombField = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#FF0000"))
	Foreground = lipgloss.Color("#FFFFFF")

	BombIcon       = "o"
	FlagIcon       = "o"
	UnrevealedIcon = "?"

	VictoryMessage = lipgloss.NewStyle().
			Width(60).
			Height(20).
			Align(lipgloss.Center, lipgloss.Center).
			Bold(true).
			Background(lipgloss.Color("#3D3634")).
			Foreground(lipgloss.Color("#00FF00"))

	DefeatMessage = lipgloss.NewStyle().
			Width(60).
			Height(20).
			Align(lipgloss.Center, lipgloss.Center).
			Bold(true).
			Background(lipgloss.Color("#3D3634")).
			Foreground(lipgloss.Color("#FF0000"))
)
