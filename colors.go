package main

type rgbColor struct {
	r int
	g int
	b int
}

var colorMap = [36]rgbColor{
	rgbColor{7, 7, 7},
	rgbColor{31, 7, 7},
	rgbColor{47, 15, 7},
	rgbColor{71, 15, 7},
	rgbColor{87, 23, 7},
	rgbColor{103, 31, 7},
	rgbColor{119, 31, 7},
	rgbColor{143, 39, 7},
	rgbColor{159, 47, 7},
	rgbColor{175, 63, 7},
	rgbColor{191, 71, 7},
	rgbColor{199, 71, 7},
	rgbColor{223, 79, 7},
	rgbColor{223, 87, 7},
	rgbColor{223, 87, 7},
	rgbColor{215, 95, 7},
	rgbColor{215, 103, 15},
	rgbColor{207, 111, 15},
	rgbColor{207, 119, 15},
	rgbColor{207, 127, 15},
	rgbColor{207, 135, 23},
	rgbColor{199, 135, 23},
	rgbColor{199, 143, 23},
	rgbColor{199, 151, 31},
	rgbColor{191, 159, 31},
	rgbColor{191, 159, 31},
	rgbColor{191, 167, 39},
	rgbColor{191, 167, 39},
	rgbColor{191, 175, 47},
	rgbColor{183, 175, 47},
	rgbColor{183, 183, 47},
	rgbColor{183, 183, 55},
	rgbColor{207, 207, 111},
	rgbColor{223, 223, 159},
	rgbColor{239, 239, 199},
	rgbColor{255, 255, 255},
}
