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

func (f *Food) Draw(parent [][]string) [][]string {
	view := parent
	parent[f.PositionY][f.PositionX] = "♡"
	return view
}
