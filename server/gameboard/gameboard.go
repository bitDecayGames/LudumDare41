package gameboard

type Tile struct {
	Passable bool
}

type GameBoard struct {
	tiles [][]Tile
}

func LoadBoard(name string) GameBoard {
	// TODO: Load this from file / config
	return GameBoard{
		tiles: [][]Tile{
			{Tile{Passable: true}},
			{Tile{Passable: true}},
		},
	}
}
