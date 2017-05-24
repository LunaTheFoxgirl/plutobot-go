package simulation

import "math"

type Vector struct {
	X float64
	Y float64
}

func DeltaV(totalmass, emptymass, velocity float64) float64 {
	return velocity * math.Log(totalmass/emptymass)
}

