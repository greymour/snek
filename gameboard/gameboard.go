package gameboard

import (
	"fmt"
	"math/rand/v2"
	"snek/food"
	"snek/snake"
	"strings"
)

type Cell struct{}

type GameBoard struct {
	Snake  *snake.Snake
	Width  int
	Height int
	Food   *food.Food
	Score  int
}

const ROW_OFFSET = 1

func (gb *GameBoard) GetWidth() int {
	return gb.Width
}

func (gb *GameBoard) GetHeight() int {
	return gb.Height
}

// right now we're just basing the bounds of the GameBoard on math calculations
// there's no underlying data structure here
func New(rowCount int, columnCount int, snake *snake.Snake) *GameBoard {
	// adding 2 to the Height to allow for the "Score" read out and borders
	return &GameBoard{Snake: snake,
		Width:  rowCount,
		Height: columnCount + 2,
		Score:  0,
	}
}

// due to using a _ and ▔ character we need to offset things by 1 to make it look right
func (gb *GameBoard) SnakeOutsideBounds() bool {
	if (gb.Snake.PositionX <= 0 || gb.Snake.PositionX >= gb.Width-1) || (gb.Snake.PositionY <= 1 || gb.Snake.PositionY >= gb.Height-1) {
		return true
	}
	return false
}

func (gb *GameBoard) SpawnFood() {
	gb.Food = nil
	x := rand.IntN(gb.Width-2) + 1
	y := rand.IntN(gb.Height-2-ROW_OFFSET) + 1 + ROW_OFFSET
	for gb.Snake.HasCellAt(x, y) {
		x = rand.IntN(gb.Width-2) + 1
		y = rand.IntN(gb.Height-2-ROW_OFFSET) + 1 + ROW_OFFSET
	}
	gb.Food = food.New(x, y)
}

func (gb *GameBoard) SpawnFoodAt(x int, y int) {
	gb.Food = food.New(x, y)
}

func (gb *GameBoard) Tick() {
	if gb.Food != nil {
		if gb.Snake.HasCellAt(gb.Food.PositionX, gb.Food.PositionY) {
			// @TODO: need to increment the life of the Food here if I want that feature
			gb.Snake.EatFood(0, gb.Width, 0, gb.Height)
			gb.Food = nil
			gb.Score++
			gb.SpawnFood()
		}
	}
}

func (gb *GameBoard) HasCellAt(x int, y int) bool {
	if x < 0 || x > gb.Width-1 || y < 0 || y > gb.Height-1 {
		return false
	}
	return true
}

func (gb *GameBoard) isCorner(x int, y int) bool {
	if (x == 0 || x == gb.Width-1) && (y == ROW_OFFSET || y == gb.Height-1) {
		return true
	}
	return false
}

func (gb *GameBoard) makeEmptyRow(y int) []string {
	emptyRow := []string{}
	for i := 0; i < gb.Width; i++ {
		char := ""
		if gb.isCorner(i, y) {
			char = " "
		} else if y == ROW_OFFSET {
			char = "_"
		} else if y == gb.Height-1 {
			char = "▔"
		} else {
			if i == 0 || i == gb.Width-1 {
				char = "|"
			} else {
				char = "·"
			}
		}
		emptyRow = append(emptyRow, char)
	}
	return emptyRow
}

func (gb *GameBoard) drawScore(parent [][]string) [][]string {
	view := parent
	scoreStr := strings.Split(fmt.Sprintf("Score: %v", gb.Score), "")
	if len(view[0]) == 0 {
		panic("\nLength of first row in view is too small to contain score string, process will exit\n")
	}
	for i, c := range scoreStr {
		view[0][i] = c
	}
	return view
}

// example of score: `Score: 0`, `Score: 100`
func (gb *GameBoard) Draw(parent [][]string, paddingX int, paddingY int) [][]string {
	view := parent
	limit := gb.Width * gb.Height
	currentRow := ROW_OFFSET
	for i := 0; i < limit; i++ {
		if i > 0 && i%gb.Width == 0 {
			view[currentRow] = gb.makeEmptyRow(currentRow)
			currentRow++
		}
	}
	view = gb.drawScore(view)
	if gb.Food != nil {
		view = gb.Food.Draw(view)
	}
	view = gb.Snake.Draw(view)
	return view
}
