package main

import "errors"

var errNegativeWidth = errors.New("negative width")
var errMinimumHeight = errors.New("height should be at least one")
var errMinimumThreshold = errors.New("threshold value should be a positive integer")
