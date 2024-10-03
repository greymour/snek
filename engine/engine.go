package engine

import (
	"snek/gameboard"
	"snek/input"
	"snek/menu"
)

type GameState int

const (
	MENU GameState = iota
	ACTIVE_GAME
	PAUSE
	LOADING
)

type Engine struct {
	menu      *menu.Menu
	gameboard *gameboard.GameBoard
	//	player      *player.Player
	inputBuffer *input.InputBuffer
}

// func New() *Engine {
// 	m := menu.Menu{}
// 	gb := gameboard.New()
// 	ib := input.New()
// 	return &Engine{}
// }

func (e *Engine) Init() {

}

func (e *Engine) Pause() {

}

func (e *Engine) UnPause() {

}

func (e *Engine) Save() {

}

func (e *Engine) Quit() {
	e.Save()
	e.inputBuffer.Close()
}

func (e *Engine) StartGame() {

}

func (e *Engine) OpenMenu() {

}
