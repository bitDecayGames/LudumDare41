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
	// Used by respawn logic, DO NOT USE unless you know what you're doing
	TempOccupied bool
}

type GameBoard struct {
	Tiles          [][]Tile
	FlattenedTiles []Tile `json:"tiles"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
}

var gameMap = [][]int{
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func (gb GameBoard) GetTilesByType(tileType string) []Tile {
	matchingTiles := []Tile{}
	for _, tile := range gb.FlattenedTiles {
		if tile.TileType == tileType {
			matchingTiles = append(matchingTiles, tile)
		}
	}
	return matchingTiles
}

func (gb GameBoard) OnBoard(pos utils.Vector) bool {
	return pos.X >= 0 && pos.X < len(gb.Tiles) &&
		pos.Y >= 0 && pos.Y < len(gb.Tiles[0])
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
	for col := 0; col < len(gameMap); col++ {
		board.Tiles = append(board.Tiles, []Tile{})
		for row := 0; row < len(gameMap[col]); row++ {
			tile := Tile{
				TileType: Empty_tile,
			}
			if gameMap[col][row] == 1 {
				tile.TileType = Wall_tile
			}
			board.Tiles[col] = append(board.Tiles[col], tile)
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
