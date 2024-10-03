package renderer

import (
	"fmt"
	"golang.org/x/term"
	"snek/utils"
)

type Drawable interface {
	// @TODO: make this also return an error if there's a problem drawing the characters - eg. if the image being drawn
	// is bigger than the parent
	Draw(parent [][]string) [][]string
	// maybe rename this to `OccupiesCell`?
	HasCellAt(int, int) bool
}

type DrawableRoot interface {
	Draw(parent [][]string, paddingX int, paddingY int) [][]string
	GetWidth() int
	GetHeight() int
}

type Renderer struct {
	View          [][]string
	EmptyCellChar rune
	BorderXChar   rune
	BorderYChar   rune
	size          int
}

func New() *Renderer {
	return &Renderer{
		View: [][]string{},
		size: 0,
	}
}

func (r *Renderer) padTop(offsetY int) {
	for i := len(r.View) - offsetY - 1; i >= 0; i-- {
		r.View[i+offsetY] = r.View[i]
		r.View[i] = []string{}
	}
}

func (r *Renderer) padLeft(offsetX int) {
	for i, row := range r.View {
		padding := make([]string, offsetX)
		utils.FillSlice(padding, " ")
		newRow := append(padding, row...)
		r.View[i] = newRow
	}
}

// @TODO: implement these util functions
func AlignLeft() {}

func AlignRight() {}

func AlignTop() {}

func AlignBottom() {}

func Center(parent [][]string, child [][]string) [][]string {
	return [][]string{}
}

func (r *Renderer) CreateViewModel(root DrawableRoot) {
	width, height, err := term.GetSize(0)
	height = height - 2
	offsetX := (width - root.GetWidth()) / 2
	offsetY := (height - root.GetHeight()) / 2

	if err != nil {
		panic("could not get size of window, rip")
	}
	r.View = make([][]string, 0, height)
	for i := 0; i < height; i++ {
		newRow := make([]string, width, width)
		utils.FillSlice(newRow, " ")
		r.View = append(r.View, newRow)
	}

	r.View = root.Draw(r.View, offsetX, offsetY)
	// why do I pad top twice? I stopped keeping track of the math a while ago
	// answer: it's because I'm appending way too many rows at some point...
	// so the first padTop shoves it down to the top of the viewport, and the second padTop
	// moves it down 50%
	// @TODO: FIX THE MATH
	r.padTop(offsetY)
	r.padTop(offsetY / 2)
	r.padLeft(offsetX)
}

func (r *Renderer) RenderView(root DrawableRoot) {
	r.CreateViewModel(root)
	view := "\n"

	for _, row := range r.View {
		for _, character := range row {
			view += character
		}
		view += "\n"
	}
	fmt.Println(view)
}
