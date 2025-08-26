package solutions

import (
	"math/rand"
	"strconv"
	"time"
)

// GenerateRandomCardNumber generates a random 16-digit card number as a string
func GenerateRandomCardNumber() string {
	rand.Seed(time.Now().UnixNano())
	num := ""
	for i := 0; i < 16; i++ {
		digit := rand.Intn(10)
		num += strconv.Itoa(digit)
	}
	return num
}
