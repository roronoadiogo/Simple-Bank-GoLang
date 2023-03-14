package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random intger between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(lengthName int) string {

	var builder strings.Builder
	character := len(alphabet)

	for i := 0; i < lengthName; i++ {
		onlyChar := alphabet[rand.Intn(character)]
		builder.WriteByte(onlyChar)
	}

	return builder.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	randomCurrency := len(currencies)
	return currencies[rand.Intn(randomCurrency)]
}
