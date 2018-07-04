package utils

import (
	"math/rand"
	"time"
)

func Random(min, max int) int {
	if min == max {
		return min
	}

	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
