package minesweeper

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/igor-mauricio/classic-games-on-terminal/minesweeper/styles"
)

func (f Game) gameView() string {
	result := ""
	for i := range f.Positions {
		for j := range f.Positions[i] {
			var style lipgloss.Style
			text := ""
			if f.Positions[i][j].isFlagged {
				style = styles.FlaggedField
				text = styles.FlagIcon
			} else if !f.Positions[i][j].isRevealed {
				style = styles.UnrevealedField
				text = styles.UnrevealedIcon
			} else if f.Positions[i][j].isBomb {
				style = styles.BombField
				text = styles.BombIcon
			} else {
				style = styles.RevealedField
				if f.Positions[i][j].nearbyBombs == 0 {
					text = " "
				} else {
					text = strconv.Itoa(f.Positions[i][j].nearbyBombs)
				}
			}
			if f.pos.row == i && f.pos.col == j {
				style = style.Background(styles.Foreground)
			}
			result += style.Padding(0, 1).Render(text)
		}
		result += "\n"
	}
	return result
}
