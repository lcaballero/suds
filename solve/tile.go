package solve

type Tile struct {
	Value int
	Excluded map[int]struct{}
}

func NewTile(v int) *Tile {
	return &Tile{
		Value: v,
		Excluded: make(map[int]struct{}),
	}
}

func (t *Tile) HasValue() bool {
	return 0 < t.Value && t.Value < 10
}

