package gameboard

const empty_tile = "empty"
const wall_tile = "wall"

type Tile struct {
	ID       int
	TileType string `json:"tileType"`
}

type GameBoard struct {
	tiles [][]Tile
}

func LoadBoard(name string) GameBoard {
	// TODO: Load this from file / config
	return GameBoard{
		tiles: [][]Tile{
			{Tile{TileType: empty_tile}},
			{Tile{TileType: empty_tile}},
		},
	}
}
