package grid

import (
	"math/rand"
	"sync"
	"time"

	"github.com/juliendoutre/doom-fire/utils"
)

// New create a new Grid object.
func New(
	width, height, initCondition int,
	threshold float64,
) (*Grid, error) {
	if height < 1 {
		return nil, &errMinimumValue{"height", 1}
	}

	if width < 0 {
		return nil, &errMinimumValue{"width", 0}
	}

	if threshold <= 0 {
		return nil, &errMinimumValue{"threshold", 0}
	}

	cells := make([][]int, height)

	// init all rows with 0.
	for i := 0; i < height; i++ {
		cells[i] = make([]int, width)
	}

	// initialize last row with init condition value.
	for j := 0; j < width; j++ {
		cells[height-1][j] = initCondition
	}

	return &Grid{
		cells:     cells,
		threshold: threshold,
		wind:      threshold,
	}, nil
}

// Grid stores the cells state and evolution parameters.
type Grid struct {
	sync.Mutex
	cells           [][]int
	threshold, wind float64
}

// Cell returns the value in a cell.
func (g *Grid) Cell(x, y int) int {
	g.Lock()
	value := g.cells[x][y]
	g.Unlock()

	return value
}

// NextState compute the next state of the grid.
func (g *Grid) NextState() {
	g.Lock()

	for y := len(g.cells) - 2; y > -1; y-- {
		for x := 0; x < len(g.cells[y]); x++ {
			r := int(rand.Float64() * g.threshold)
			w := rand.Float64() * g.wind
			s := r
			if w > .5 {
				s = -s
			}
			d := utils.BoundInt(x+s, 0, len(g.cells[y])-1)
			g.cells[y][x] = utils.Max(g.cells[y+1][d]-r, 0)
		}
	}

	g.Unlock()
}

// Start runs a simulation with a certain interval of time between each state computation.
func (g *Grid) Start(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			g.NextState()
		}
	}
}

// UpdateThreshold increases or decreases the threshold value.
func (g *Grid) UpdateThreshold(v float64) {
	g.Lock()
	g.threshold += v
	g.Unlock()
}

// UpdateWind increases or decreases the wind value.
func (g *Grid) UpdateWind(v float64) {
	g.Lock()
	g.wind += v
	g.Unlock()
}
