package main

import (
	"sync"

	tl "github.com/JoelOtter/termloop"
)

type point struct {
	*tl.Rectangle
	m     *sync.Mutex
	x, y  int
	cells [][]int
}

func (p *point) Tick(el tl.Event) {
	//p.m.Lock()
	p.SetColor(tl.RgbTo256Color(colorMap[p.cells[p.y][p.x]].r, colorMap[p.cells[p.y][p.x]].g, colorMap[p.cells[p.y][p.x]].b))
	//p.m.Unlock()
}
