package crypto

import (
	"fmt"
	"math/rand"
)

const encodeURL = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var (
	randStrSpaceLen = len(encodeURL)
)

func GenRandomNumber(length int) string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func GenRandomString(length int) string {
	var randStr string
	for i := 0; i < length; i++ {
		randStr += string(encodeURL[rand.Intn(randStrSpaceLen)])
	}
	return randStr
}
