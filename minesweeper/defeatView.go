package minesweeper

import (
	"github.com/igor-mauricio/classic-games-on-terminal/minesweeper/styles"
)

func DefeatView() string {
	result := ""
	result += styles.DefeatMessage.Render("Defeat\nPress enter to retry")
	return result
}
