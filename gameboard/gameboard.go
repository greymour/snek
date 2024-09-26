package gameboard

import (
	"snek/food"
	"snek/renderer/rendercell"
	"snek/snake"
)

type Cell struct{}

type GameBoard struct {
	Snake  *snake.Snake
	Width  int
	Height int
	Food   *food.Food
	// Cells [][]Cell
}

// right now we're just basing the bounds of the GameBoard on math calculations
// there's no underlying data structure here
func New(rowCount int, columnCount int, snake *snake.Snake) *GameBoard {
	// this is wrong but eh
	// columns := []Cell{}
	// rows := []Cell{}
	// cells := []Cell{}
	// numRows := 0
	// // I need to check if a list of columns exists at the current i idx
	// // if not and i % columnCount is zero, we need to push a new list into cells
	// // after that I need to push a new Cell
	// for i := 0; i < rowCount; i++ {
	// 	// need to append a new slice, start appending at that idx
	// 	if i%columnCount == 0 {

	// 	}
	// 	rows = append(rows, Cell{})
	// 	numRows++
	// }
	return &GameBoard{Snake: snake, Width: rowCount, Height: columnCount}
}

func (gb *GameBoard) SnakeOutsideBounds(snake *snake.Snake) bool {
	if (snake.PositionX <= 0 || snake.PositionX >= gb.Width) || (snake.PositionY <= 0 || snake.PositionY >= gb.Height) {
		return true
	}
	return false
}

func (gb *GameBoard) SpawnFood() {
	// @TODO: spawn these randomly in a location not occupied by the snake and within the bounds of the board
	x := 5
	y := 10
	gb.Food = food.New(x, y)
}

func (gb *GameBoard) SpawnFoodAt(x int, y int) {
	gb.Food = food.New(x, y)
}

func (gb *GameBoard) DespawnFood() {
	gb.Food = nil
}

// do everything that we need to advance the game on tick of the program
// @TODO: right now this doesn't handle user input at all
// @TODO: this shouldn't live here I think maybe? dunno!
// could put this in an Engine or Runtime object or s o m e t h i n g
func (gb *GameBoard) Tick() {
	if gb.Food != nil {
		// @TODO: need to increment the life of the Food here
		if (gb.Food.PositionX == gb.Snake.PositionX) && (gb.Food.PositionY == gb.Snake.PositionY) {
			gb.Snake.EatFood(0, gb.Width, 0, gb.Height)
			gb.Food = nil
		}
	}
}

func (gb *GameBoard) Draw(x int, y int) rendercell.RenderCell {
	if gb.Food != nil {
		if gb.Food.PositionX == x && gb.Food.PositionY == y {
			return "♡"
		}
	}
	// @TODO: this doesn't take into account when we have overlapping characters with the snake I guess
	if x == 0 || x == gb.Width-1 {
		return "|"
	} else if y == 0 {
		return "▔"
	} else if y == gb.Height-1 {
		return "_"
	}
	return "·"
}
