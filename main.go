package main

import (
	"fmt"
	"snek/gameboard"
	"snek/input"
	"snek/renderer"
	"snek/snake"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	inputChan := make(chan string)
	syncChan := make(chan int)
	exitChan := make(chan int)
	ib := input.New(inputChan, syncChan, exitChan)
	snake := snake.New(3, 5, 5)
	gb := gameboard.New(15, 15, snake)
	gb.SpawnFood()
	r := renderer.New()

	fmt.Println("starting listen")
	ib.OnInput(input.InputHandler{
		Key: keyboard.KeyArrowDown,
		Handler: func() {
			fmt.Printf("MOVING DOWNNN\n")
			snake.Move(input.DOWN)
		}})

	ib.OnInput(input.InputHandler{
		Key: keyboard.KeyArrowUp,
		Handler: func() {
			snake.Move(input.UP)
		}})

	ib.OnInput(input.InputHandler{
		Key: keyboard.KeyArrowRight,
		Handler: func() {
			snake.Move(input.RIGHT)
		}})

	ib.OnInput(input.InputHandler{
		Key: keyboard.KeyArrowLeft,
		Handler: func() {
			snake.Move(input.LEFT)
		}})

	ib.OnInput(input.InputHandler{
		Key: keyboard.KeyCtrlC,
		Handler: func() {
			exitChan <- 0
		},
	})

	quit := func() {
		ib.Close()
		close(inputChan)
		close(syncChan)
		close(exitChan)
		fmt.Printf("Good game! Your score: %d", gb.Score)
	}

	go ib.Listen()

	tickRate := 1 / (gb.Score + 1)

	for range time.Tick(time.Duration(tickRate) * time.Second) {
		if gb.SnakeOutsideBounds() || snake.CollisionWithSelf() {
			quit()
			return
		}
		switch ib.LastInputKeyCode {
		case keyboard.KeyCtrlC:
			quit()
			return
		case keyboard.KeyArrowLeft:
			snake.Move(input.LEFT)
		case keyboard.KeyArrowUp:
			snake.Move(input.UP)
		case keyboard.KeyArrowRight:
			snake.Move(input.RIGHT)
		case keyboard.KeyArrowDown:
			snake.Move(input.DOWN)
		}
		gb.Tick()
		r.RenderView(gb)
	}
}
