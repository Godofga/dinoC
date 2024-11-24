package minesweeper

func (g *Game) checkStatus() gameStatus {
	stillPlaying := false

	for i := range g.Positions {
		for j := range g.Positions[i] {
			if g.Positions[i][j].isBomb && g.Positions[i][j].isRevealed {
				return DEFEAT
			}
			if g.Positions[i][j].isBomb && !g.Positions[i][j].isFlagged ||
				!g.Positions[i][j].isBomb && !g.Positions[i][j].isRevealed {
				stillPlaying = true
			}
		}
	}
	if stillPlaying {
		return PLAYING
	}
	return VICTORY
}
