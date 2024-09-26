package snake

import (
	"fmt"
	"snek/input"
	"snek/renderer/rendercell"
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
	lastMovedDirection input.Direction
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
		lastMovedDirection: input.DOWN,
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

func (s *Snake) Move(d input.Direction) {
	lastMove := s.lastMovedDirection
	moved := false
	s.lastMovedDirection = d
	switch s.lastMovedDirection {
	case input.UP:
		if lastMove != input.DOWN {
			s.PositionY -= 1
			moved = true
		}
	case input.DOWN:
		if lastMove != input.UP {
			s.PositionY += 1
			moved = true
		}
	case input.LEFT:
		if lastMove != input.RIGHT {
			s.PositionX -= 1
			moved = true
		}
	case input.RIGHT:
		if lastMove != input.LEFT {
			s.PositionX += 1
			moved = true
		}
	default:
		panic("Received Direction argument that was not UP/DOWN/LEFT/RIGHT")
	}

	if !moved {
		return
	}

	for i := s.size - 1; i >= 0; i-- {
		// we could just start at 1, but leaving this here in case there's more logic we want to do on each
		// segment
		seg := s.segments[i]
		// fmt.Printf("seg: %v", seg)
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
			case input.UP:
				return "▲"
			case input.DOWN:
				return "▼"
			case input.RIGHT:
				return "▶"
			case input.LEFT:
				return "◀"
			default:
				return "▼"
			}
		}
	}

	// return "□"
	return "■"
}
