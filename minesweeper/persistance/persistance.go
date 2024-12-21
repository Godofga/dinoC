package persistance

import "fmt"

type Database struct {
	positions [][]Cell
	status    gameStatus
	pos       struct{ row, col int }
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) CreateBoard(rows, cols int) error {
	if rows < 0 || cols < 0 {
		return fmt.Errorf("out of bounds")
	}
	d.positions = make([][]Cell, rows)
	for i := range d.positions {
		d.positions[i] = make([]Cell, cols)
	}
	return nil
}

func (d *Database) GetPositions() [][]Cell {
	return d.positions
}

func (d *Database) UpdatePositions(pos [][]Cell) error {
	if len(pos) != len(d.positions) {
		return fmt.Errorf("invalid row size")
	}
	if len(pos[0]) != len(d.positions[0]) {
		return fmt.Errorf("invalid column size")
	}
	d.positions = pos
	return nil
}

func (d *Database) GetSize() (int, int) {
	return len(d.positions), len(d.positions[0])
}

func (d *Database) GetStatus() gameStatus {
	return d.status
}

func (d *Database) SetStatus(status gameStatus) error {
	if status != PLAYING && status != VICTORY && status != DEFEAT {
		return fmt.Errorf("invalid status")
	}
	d.status = status
	return nil
}

func (d *Database) GetPos() (int, int) {
	return d.pos.row, d.pos.col
}

func (d *Database) SetPos(row, col int) error {
	if row < 0 || row >= len(d.positions) {
		return fmt.Errorf("invalid row")
	}
	if col < 0 || col >= len(d.positions[0]) {
		return fmt.Errorf("invalid column")
	}
	d.pos.row = row
	d.pos.col = col
	return nil
}

type Cell struct {
	NearbyBombs int
	IsBomb      bool
	IsRevealed  bool
	IsFlagged   bool
}

type gameStatus int

const (
	PLAYING gameStatus = iota
	VICTORY
	DEFEAT
)

type Difficulty int

const (
	EASY Difficulty = iota
	MEDIUM
	HARD
)
