package solve

import (
	"strconv"
	"strings"
)

const (
	COLS = 9
	ROWS = 9
	NoValue = 0
)


type Board struct {
	Tiles [][]*Tile
}

func NewEmptyBoard() *Board {
	tiles := make([][]*Tile, ROWS)
	for y := 0; y < ROWS; y++ {
		tiles[y] = make([]*Tile, COLS)
		for x := 0; x < COLS; x++ {
			tiles[y][x] = NewTile(NoValue)
		}
	}

	return &Board{
		Tiles: tiles,
	}
}

func NewBoard(vals [][]string) (*Board, error) {
	tiles := make([][]*Tile, ROWS)
	for r,row := range vals {
		tiles[r] = make([]*Tile, COLS)
		for c,item := range row {
			s := strings.Trim(item, " \t\n\r")
			v, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			tiles[r][c] = NewTile(v)
		}
	}
	board := &Board{
		Tiles:tiles,
	}

	return board, nil
}

func (b *Board) PlaysLeft() int {
	withVals := 0
	for x := 0; x < ROWS; x++{
		for y := 0; y < COLS; y++ {
			t := b.Tiles[x][y]
			if t.HasValue() {
				withVals++
			}
		}
	}
	return (COLS * ROWS) - withVals
}

