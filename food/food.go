package food

type Food struct {
	PositionX       int
	PositionY       int
	MaxLifetime     int
	CurrentLifetime int
}

func New(x int, y int) *Food {
	return &Food{x, y, 5, 0}
}
