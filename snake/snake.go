package snake

import (
	"fmt"
	"snek/renderer/rendercell"
)

type Direction string

const (
	START Direction = "START"
	LEFT  Direction = "LEFT"
	RIGHT Direction = "RIGHT"
	UP    Direction = "UP"
	DOWN  Direction = "DOWN"
)

type SnakeSegment struct {
	x int
	y int
}

type Snake struct {
	size int
	// these two positions are the position of the snake's head
	PositionX          int
	PositionY          int
	lastMovedDirection Direction
	segments           []*SnakeSegment
}

// startX and startY are the coordinates of the snake's head
func New(startingSize int, startX int, startY int) *Snake {
	if startingSize < 1 {
		panic("received starting size of less than 1 for snake, process will exit")
	}
	segments := []*SnakeSegment{}
	for i := 0; i < startingSize; i++ {
		segments = append(segments, &SnakeSegment{startX, startY - i})
	}
	return &Snake{
		size:               startingSize,
		PositionX:          startX,
		PositionY:          startY,
		segments:           segments,
		lastMovedDirection: START,
	}
}

func (s *Snake) finishMove() {
	// iterate backwards over the segments in the Snake, update their coords
	for i := s.size - 1; i >= 0; i-- {
		// we could just start at 1, but leaving this here in case there's more logic we want to do on each
		// segment
		seg := s.segments[i]
		fmt.Printf("seg: %v", seg)
		if i == 0 {
			seg.x = s.PositionX
			seg.y = s.PositionY
		} else {
			// if we index outside the slice, we get a runtime exception :D
			prev := s.segments[i-1]
			seg.x = prev.x
			seg.y = prev.y
		}
	}
}

func (s *Snake) CollisionWithSelf() bool {
	// due to the snake only moving at right angles, it's impossible for it to collide with itself unless it has
	// a size of at least 5
	if s.size < 5 {
		return false
	}
	// @TODO: as an optimization for very long Snakes, iterate backwards based on some relative calculation?
	// eg. compare the x and y values of the first and last Segments in the snake, and do something based on that?
	// something to think about, although this optimization definitely doesn't matter - fun to think about though
	head := s.segments[0]
	for i := 4; i < len(s.segments); i++ {
		if head.x == s.segments[i].x && head.y == s.segments[i].y {
			return true
		}
	}
	return false
}

// func move(x int, y int) {

// }

// @TODO: Combine these into one method and instead take in a Direction as an argument instead of ints
func (s *Snake) MoveX(n int) {
	lastMove := s.lastMovedDirection
	originalX := s.PositionX
	if n < 0 {
		s.lastMovedDirection = LEFT
	} else if n > 0 {
		s.lastMovedDirection = RIGHT
	} else {
		panic("Snake did not move, crashing my shit")
	}
	if lastMove != START && (lastMove == LEFT || lastMove == RIGHT) {
		if lastMove == s.lastMovedDirection {
			s.PositionX += n
		} else {
			fmt.Println("Not moving Snake X")
			panic("BUGGG")
		}
	} else {
		s.PositionX += n
	}
	fmt.Printf("Original and new position X: %v %v", lastMove, s.lastMovedDirection)
	if originalX != s.PositionX {
		s.finishMove()
	}
}

func (s *Snake) MoveY(n int) {
	lastMove := s.lastMovedDirection
	originalY := s.PositionY
	if n < 0 {
		s.lastMovedDirection = UP
	} else if n > 0 {
		s.lastMovedDirection = DOWN
	} else {
		panic("Snake did not move, crashing my shit")
	}
	if lastMove != START && (lastMove == UP || lastMove == DOWN) {
		if lastMove == s.lastMovedDirection {
			s.PositionY += n
		} else {
			fmt.Println("Not moving Snake Y")
			panic("BUGGG")
		}
	} else {
		s.PositionY += n
	}

	fmt.Printf("Original and new position Y: %v %v", lastMove, s.lastMovedDirection)
	if originalY != s.PositionY {
		s.finishMove()
	}
}

func (s *Snake) HasCellAt(x int, y int) bool {
	for i := 0; i < len(s.segments); i++ {
		seg := s.segments[i]
		if seg.x == x && seg.y == y {
			return true
		}
	}
	return false
}

// @TODO: need to know where the borders of the gameboard are here for where to put the tail
func (s *Snake) EatFood(borderXStart int, borderXEnd int, borderYStart int, borderYEnd int) {
	fmt.Printf("snake eating fooood")
	lastSegment := s.segments[len(s.segments)-1]
	x := lastSegment.x
	y := lastSegment.y
	if lastSegment.x <= 0 {
		x = lastSegment.x + 1
	} else if lastSegment.x >= borderXEnd {
		x = lastSegment.x - 1
	} else if lastSegment.y <= borderYStart {
		y = lastSegment.y + 1
	} else if lastSegment.y >= borderYEnd {
		y = lastSegment.y - 1
	}

	s.segments = append(s.segments, &SnakeSegment{x, y})
	s.size++
}

func (s *Snake) PrintSegments() {
	for i := 0; i < len(s.segments); i++ {
		fmt.Printf("segment: %+v", s.segments[i])
	}
}

func (s *Snake) Draw(x int, y int) rendercell.RenderCell {
	if s.PositionX == x && s.PositionY == y {
		if s.CollisionWithSelf() {
			return "▣"
		} else {
			switch s.lastMovedDirection {
			case UP:
				return "▲"
			case DOWN:
				return "▼"
			case RIGHT:
				return "▶"
			case LEFT:
				return "◀"
			default:
				return "▼"
			}
		}
	}

	// return "□"
	return "■"
}
