package service

import (
	"github.com/igor-mauricio/classic-games-on-terminal/minesweeper/persistance"
)

func (s *Service) NewGame(rows, cols int) error {
	s.db.CreateBoard(rows, cols)
	s.db.SetStatus(persistance.PLAYING)

	return nil
}
