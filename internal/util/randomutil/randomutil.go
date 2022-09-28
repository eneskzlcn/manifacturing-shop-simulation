package randomutil

import (
	"math/rand"
	"time"
)

func RandomInt(lowerBound, upperBound int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(upperBound-lowerBound) + lowerBound
}
