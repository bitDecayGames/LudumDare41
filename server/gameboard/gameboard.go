package gameboard

import (
	"github.com/bitDecayGames/LudumDare41/server/utils"
)

const (
	Empty_tile = "empty"
	Wall_tile  = "wall"

	width  = 5
	height = 6
)

type Tile struct {
	ID       int          `json:"id"`
	TileType string       `json:"tileType"`
	Pos      utils.Vector `json:"pos"`
}

type GameBoard struct {
	Tiles          [][]Tile
	FlattenedTiles []Tile `json:"tiles"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
}

func LoadBoard(name string) GameBoard {
	// TODO: Load this from file / config
	board := GameBoard{
		Tiles:          [][]Tile{},
		FlattenedTiles: []Tile{},
		Width:          width,
		Height:         height,
	}

	// Generate tiles
	for x := 0; x < width; x++ {
		board.Tiles = append(board.Tiles, []Tile{})
		for y := 0; y < height; y++ {
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
