package solve

import (
	"bytes"
	"encoding/csv"
)


func tiles(csvtext string) [][]string {
	reader := csv.NewReader(bytes.NewReader([]byte(csvtext)))
	reader.TrimLeadingSpace = true
	lines, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return lines
}

func board(v string) (*Board, error) {
	return NewBoard(tiles(v))
}


func board1() (*Board,error) {
	return board(`
0,0,0, 0,0,0, 0,0,0
0,0,0, 0,0,0, 0,0,0
0,0,0, 0,0,0, 0,0,0

0,0,0, 0,0,0, 0,0,0
0,0,0, 0,0,0, 0,0,0
0,0,0, 0,0,0, 0,0,0

0,0,0, 0,0,0, 0,0,0
0,0,0, 0,0,0, 0,0,0
0,0,0, 0,0,0, 0,0,0
`)
}


func board2() (*Board,error) {
	return board(`
1,0,9, 0,2,5, 3,8,0
8,6,0, 0,0,3, 0,0,0
0,0,7, 0,0,0, 5,0,9

0,0,8, 4,0,0, 6,0,0
0,1,0, 2,9,7, 0,4,0
0,0,4, 0,0,6, 1,0,0

3,0,6, 0,0,0, 4,0,0
0,0,0, 3,0,0, 0,6,2
0,5,2, 8,6,0, 9,0,1
`)
}

