package main

import (
	"time"

	tl "github.com/JoelOtter/termloop"
	"github.com/juliendoutre/doom-fire/grid"
)

func main() {
	width, height, threshold := parseArgs()

	gr, err := grid.New(width, height, len(colorMap)-1, threshold)
	if err != nil {
		panic(err)
	}

	g := tl.NewGame()
	l := tl.NewBaseLevel(tl.Cell{})

	g.Screen().SetLevel(l)
	g.Screen().EnablePixelMode()

	l.AddEntity(&eventCatcher{
		tl.NewRectangle(0, 0, 0, 0, tl.RgbTo256Color(0, 0, 0)),
		gr,
	})

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			color := colorMap[gr.Cell(j, i)]
			l.AddEntity(
				&point{
					tl.NewRectangle(
						i, j, 1, 1,
						tl.RgbTo256Color(
							color.r,
							color.g,
							color.b,
						),
					),
					i, j, gr,
				},
			)
		}
	}

	go gr.Start(166 * time.Millisecond)

	g.Start()
}
