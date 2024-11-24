package minesweeper

import (
	"github.com/igor-mauricio/classic-games-on-terminal/minesweeper/styles"
)

func VictoryView() string {
	result := ""
	result += styles.VictoryMessage.Render("Victory")
	return result
}
