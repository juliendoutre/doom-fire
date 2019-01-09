package main

import (
	tl "github.com/JoelOtter/termloop"
)

func main() {
	width, height, threshold := parseArgs()
	checkThreshold(threshold)

	cells, err := initCells(width, height)
	checkError(err)

	computeNextFrame(cells, threshold)

	g := tl.NewGame()
	l := tl.NewBaseLevel(tl.Cell{})

	g.Screen().SetLevel(l)
	g.Screen().EnablePixelMode()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			l.AddEntity(tl.NewRectangle(i, j, 1, 1, tl.RgbTo256Color(colorMap[cells[j][i]].r, colorMap[cells[j][i]].g, colorMap[cells[j][i]].b)))
		}
	}

	g.Start()
}
