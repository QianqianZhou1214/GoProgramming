package clockface

import (
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

func SecondHandInRadians(t time.Time) float64 {
	return math.Pi
}
