package main

import (
	"fmt"
	"snek/gameboard"
	"snek/renderer"
	"snek/snake"
)

func main() {
	rowCount := 10
	columnCount := 10
	snake := snake.New(3, rowCount/2, columnCount/2)
	gb := gameboard.New(10, 10, snake)
	fmt.Printf("%v %v %v\n", gb.Width, gb.Height, gb.Snake)
	r := renderer.New('·', '|', '_')
	r.RenderView(gb, snake)
	gb.SpawnFoodAt(5, 6)
	r.RenderView(gb, snake)
	snake.MoveY(-1)
	r.RenderView(gb, snake)
	snake.MoveX(-1)
	r.RenderView(gb, snake)
	snake.MoveY(-1)
	r.RenderView(gb, snake)
	snake.MoveX(1)
	r.RenderView(gb, snake)
	if snake.CollisionWithSelf() {
		fmt.Println("Snake collided with itself!")
		fmt.Printf("%+v\n", snake)
	}

	// r.RenderView(gb, snake)
	// i := 0
	// for !gb.SnakeOutsideBounds(snake) && !snake.CollisionWithSelf() {
	// 	switch i {
	// 	case 2:
	// 		snake.MoveX(-1)
	// 		break
	// 	case 4:
	// 		snake.MoveX(1)
	// 		break
	// 	default:
	// 		snake.MoveY(1)
	// 		break
	// 	}
	// 	gb.Tick()
	// 	fmt.Printf("%+v\n", gb.Snake)
	// 	fmt.Println("Moving the snake!")
	// 	snake.PrintSegments()
	// 	i++
	// 	r.RenderView(gb, snake)
	// }
	fmt.Println("It's joever!")
}

// structure
// a game board of a given size
// the game board is made up of cells
// a tick rate for how often the ui will update
// the snake, controlled by the user
// food that randomly spawns in a cell unoccupied by the snake
// need to detect collisions between the edges of the board and the snake
// need to detect collisions between the snake and itself
// need to detect collisions between the snake and food
// game board is a 2 dimensional array of x and y values
