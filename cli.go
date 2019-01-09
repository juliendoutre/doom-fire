package main

import "flag"

func parseArgs() (int, int, float64) {
	w := flag.Int("w", 100, "the canvas width")
	h := flag.Int("h", 37, "the canvas height")
	t := flag.Float64("t", 4.5, "threshold for fire propagation")
	flag.Parse()
	return *w, *h, *t
}
