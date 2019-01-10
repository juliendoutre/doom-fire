package main

import (
	"sync"

	tl "github.com/JoelOtter/termloop"
)

func main() {
	width, height, threshold := parseArgs()
	checkThreshold(threshold)

	cells, err := initCells(width, height)
	checkError(err)

	m := &sync.Mutex{}

	done := make(chan bool)

	g := tl.NewGame()
	l := tl.NewBaseLevel(tl.Cell{})

	g.Screen().SetLevel(l)
	g.Screen().EnablePixelMode()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			l.AddEntity(&point{
				tl.NewRectangle(i, j, 1, 1, tl.RgbTo256Color(colorMap[cells[j][i]].r, colorMap[cells[j][i]].g, colorMap[cells[j][i]].b)),
				m, i, j, cells,
			})
		}
	}

	go startRefresh(166, cells, threshold, done, m)

	g.Start()

	go stopRefresh(done)
}
