package main

import (
	"fmt"
	"math"
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

	fmt.Println("Welcome to Snek! Starting game...")
	fmt.Println("Press Left / Down / Right to start. Move the snake backwards to pause at any time. Press Ctrl+c to quit.")

	quit := func() {
		ib.Close()
		close(inputChan)
		close(syncChan)
		close(exitChan)
		fmt.Printf("Good game! Your score: %d", gb.Score)
	}

	go ib.Listen()

	setTickRate := func() time.Duration {
		var convertedScore float64
		if gb.Score < 1 {
			convertedScore = 1.0
		} else {
			convertedScore = float64(gb.Score)
		}
		rate := 1.0 / ((convertedScore + 1.0) + math.Log(convertedScore)) * 1000
		return time.Duration(rate) * time.Millisecond
	}

	// time.Tick() doesn't work here because the ticker variable gets cached in the for loop
	// using NewTicker() and Reset() breaks the cache and allows for the variable speed
	ticker := time.NewTicker(setTickRate())
	for range ticker.C {
		score := gb.Score
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
		if score != gb.Score {
			ticker.Reset(setTickRate())
		}
	}
}
