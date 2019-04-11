package util

import (
	"math/rand"
	"time"
)

func ShuffleIntSlice(slice []int)[]int{
	rand := rand.New(rand.NewSource(time.Now().Unix()))

	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}