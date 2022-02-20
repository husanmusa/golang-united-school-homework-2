package solution

import "math"

type sides int

const (
	SidesTriangle = sides(3)
	SidesSquare   = sides(4)
	SidesCircle   = sides(0)
	Pi            = 3.14
)

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	if sidesNum == SidesSquare {
		return sideLen * sideLen
	}
	if sidesNum == SidesTriangle {
		return math.Sqrt(3) / 4.0 * sideLen * sideLen
	}
	if sidesNum == SidesCircle {
		return Pi * sideLen * sideLen
	}

	return 0
}
