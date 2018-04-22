package gameboard

import (
	"github.com/bitDecayGames/LudumDare41/server/utils"
)

const Empty_tile = "empty"
const Wall_tile = "wall"

type Tile struct {
	ID       int          `json:"id"`
	TileType string       `json:"tileType"`
	Pos      utils.Vector `json:"pos"`
}

type GameBoard struct {
	Tiles          [][]Tile
	FlattenedTiles []Tile `json:"tiles"`
}

func LoadBoard(name string) GameBoard {
	// TODO: Load this from file / config
	board := GameBoard{
		Tiles:          [][]Tile{},
		FlattenedTiles: []Tile{},
	}

	dim := 5
	for x := 0; x < dim; x++ {
		board.Tiles = append(board.Tiles, []Tile{})
		for y := 0; y < dim; y++ {
			tile := Tile{
				TileType: Empty_tile,
			}
			board.Tiles[x] = append(board.Tiles[x], tile)
		}
	}

	// Assign x/y values
	for x, yTiles := range board.Tiles {
		for y, tile := range yTiles {
			tile.Pos = utils.Vector{
				X: x,
				Y: y,
			}
			board.Tiles[x][y] = tile
			// Flatten tiles
			board.FlattenedTiles = append(board.FlattenedTiles, tile)
		}
	}

	return board
}
