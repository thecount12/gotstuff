package helpers

import (
	"math/rand"
	"time"
)

// 3. random number
func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}

