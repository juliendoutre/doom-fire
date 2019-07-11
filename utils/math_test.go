package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	t.Run("with a > b", func(t *testing.T) {
		assert.Equal(t, 1, Max(1, 0))
	})

	t.Run("with a < b", func(t *testing.T) {
		assert.Equal(t, 1, Max(0, 1))
	})
}

func TestBounInt(t *testing.T) {
	t.Run("with x in [a, b]", func(t *testing.T) {
		assert.Equal(t, 1, BoundInt(1, 0, 1))
	})

	t.Run("with x > [a, b]", func(t *testing.T) {
		assert.Equal(t, 0, BoundInt(-1, 0, 1))
	})

	t.Run("with x < [a, b]", func(t *testing.T) {
		assert.Equal(t, 1, BoundInt(2, 0, 1))
	})
}

func TestBounFloat(t *testing.T) {
	t.Run("with x in [a, b]", func(t *testing.T) {
		assert.Equal(t, 1., BoundFloat(1., 0., 1.))
	})

	t.Run("with x > [a, b]", func(t *testing.T) {
		assert.Equal(t, 0., BoundFloat(-1., 0., 1.))
	})

	t.Run("with x < [a, b]", func(t *testing.T) {
		assert.Equal(t, 1., BoundFloat(2., 0., 1.))
	})
}
