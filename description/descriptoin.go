package description

import "errors"

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	err := errors.New("ERROR: отрицательное число")

	if height < 0 || width < 0 {
		return nil, err
	}

	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	world := &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
	return world, nil
}
