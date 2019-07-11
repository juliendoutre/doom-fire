package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGrid(t *testing.T) {
	t.Run("with invalid height", func(t *testing.T) {
		g, err := New(0, 0, 0, .1)
		assert.EqualError(t, err, (&errMinimumValue{"height", 1}).Error())
		assert.Nil(t, g)
	})

	t.Run("with negative width", func(t *testing.T) {
		g, err := New(-1, 1, 0, .1)
		assert.EqualError(t, err, (&errMinimumValue{"width", 0}).Error())
		assert.Nil(t, g)
	})

	t.Run("with negative threshold", func(t *testing.T) {
		g, err := New(0, 1, 0, .0)
		assert.EqualError(t, err, (&errMinimumValue{"threshold", 0}).Error())
		assert.Nil(t, g)
	})

	t.Run("with valid parameters", func(t *testing.T) {
		g, err := New(2, 2, 4, .1)
		assert.NoError(t, err)
		assert.Equal(t, &Grid{
			cells: [][]int{
				[]int{0, 0},
				[]int{4, 4},
			},
			threshold: .1,
			wind:      .1,
		}, g)
	})
}

func TestCell(t *testing.T) {
	g := &Grid{
		cells: [][]int{
			[]int{0, 0},
			[]int{4, 4},
		},
	}

	assert.Equal(t, 4, g.Cell(1, 0))
}

func TestNextState(t *testing.T) {
	g := &Grid{
		cells: [][]int{
			[]int{0, 0},
			[]int{4, 4},
		},
	}

	g.NextState()

	assert.Equal(t, [][]int{
		[]int{4, 4},
		[]int{4, 4},
	}, g.cells)
}

func TestUpdateThreshold(t *testing.T) {
	g := &Grid{}
	g.UpdateThreshold(1.)
	assert.Equal(t, 1., g.threshold)
}
