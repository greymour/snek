package renderer

import (
	"fmt"
	"snek/gameboard"
	"snek/renderer/rendercell"
	"snek/snake"
)

type RenderCellOLD struct {
	PositionX int
	PositionY int
	Content   rune
}

// func makeCell() *RenderCell {
// 	return &RenderCell{0, 0, '~'}
// }

type Drawable interface {
	Draw(int, int) rendercell.RenderCell
	// maybe rename this to `OccupiesCell`?
	HasCellAt(int, int) bool
}

type Renderer struct {
	View          [][]rendercell.RenderCell
	EmptyCellChar rune
	BorderXChar   rune
	BorderYChar   rune
	size          int
	numRows       int
}

func New(emptyCell rune, borderX rune, borderY rune) *Renderer {
	return &Renderer{
		View:          [][]rendercell.RenderCell{},
		EmptyCellChar: emptyCell,
		BorderXChar:   borderX,
		BorderYChar:   borderY,
		size:          0,
		numRows:       0,
	}
}

// @TODO: this is suuuuper coupled to the GameBoard, most of this logic should get moved into
// that struct's Draw method probably
func (r *Renderer) CreateViewModel(gb *gameboard.GameBoard, s *snake.Snake) {
	r.View = [][]rendercell.RenderCell{{}}
	r.numRows = 0
	r.size = 0
	currentRow := 0
	currentCol := 0
	// width := gb.Width + 1
	// height := gb.Height + 1
	limit := gb.Width * gb.Height
	for i := 0; i < limit; i++ {
		// push in a new slice to start appending to wowie
		if i > 0 && i%gb.Width == 0 {
			r.View = append(r.View, []rendercell.RenderCell{})
			currentCol = 0
			currentRow++
			r.numRows++
		}
		var cell rendercell.RenderCell
		// here is where we decide what to draw, the character for an empty space or the snake!
		if s.HasCellAt(currentCol, currentRow) {
			cell = s.Draw(currentCol, currentRow)
		} else {
			cell = gb.Draw(currentCol, currentRow)
		}
		// fmt.Printf("%v %v %v", currentRow, r.View, cell)
		r.View[currentRow] = append(r.View[currentRow], cell)
		currentCol++
		r.size++
	}
	// fmt.Printf("%+v", r)
}

func (r *Renderer) RenderView(gb *gameboard.GameBoard, s *snake.Snake) {
	view := "\n"
	r.CreateViewModel(gb, s)
	for i := 0; i <= r.numRows; i++ {
		for j := 0; j < len(r.View[i]); j++ {
			view += string(r.View[i][j])
		}
		view += "\n"
	}
	fmt.Println(view)
}
