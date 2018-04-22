package gameboard

const Empty_tile = "empty"
const Wall_tile = "wall"

type Tile struct {
	ID       int
	TileType string `json:"tileType"`
}

type GameBoard struct {
	Tiles [][]Tile
}

func LoadBoard(name string) GameBoard {
	// TODO: Load this from file / config
	return GameBoard{
		Tiles: [][]Tile{
			{Tile{TileType: Empty_tile}},
			{Tile{TileType: Empty_tile}},
		},
	}
}
