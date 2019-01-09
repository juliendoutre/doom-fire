package main

import (
	"math/rand"
)

func initCells(width, height int) (cells [][]int, err error) {
	if !checkIntMinimumValue(width, -1) {
		err = errNegativeWidth
	} else if !checkIntMinimumValue(height, 0) {
		err = errMinimumHeight
	} else {
		for y := 0; y < height-1; y++ {
			cells = append(cells, createFillArray(width, 0))
		}
		cells = append(cells, createFillArray(width, len(colorMap)-1))
	}
	return
}

func computeNextFrame(cells [][]int, threshold float64) {
	for y := len(cells) - 2; y > -1; y-- {
		for x := 0; x < len(cells[y]); x++ {
			r := int(rand.Float64() * threshold)
			d := boundInt(x+r, 0, len(cells[y])-1)
			cells[y][x] = max(cells[y+1][d]-r, 0)
		}
	}
}

func checkThreshold(threshold float64) {
	if !checkFloatMinimumValue(threshold, -1) {
		panic(errMinimumThreshold)
	}
}
