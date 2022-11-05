package random

import "math/rand"

func RangeFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RangeInt(min, max int) int {
	if min == max {
		return min
	}
	return min + rand.Intn(max-min)
}

func Chance(percent float64) bool {
	return rand.Float64() > percent
}
