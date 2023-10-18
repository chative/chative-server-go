package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

const (
	magicHashSalt = "ViEBfWhVNcDaFetC6hRx9eWXlKI76qEzgQxXG9pjFPRsLMI8BrONoW9hKptlXd3g"
)

func Sha256String(s string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(s))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func XorEncrypt(input, key string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}
	return output
}

func ServerHash(hashC, saltS string) (string, error) {
	return Sha256String(hashC + XorEncrypt(saltS, hashC) + saltS + magicHashSalt)
}

func HashID(id, saltC, saltS string) (string, error) {
	hashC, err := Sha256String(saltC + strings.ToLower(id))

	if err != nil {
		return "", err
	}
	return ServerHash(hashC, saltS)
}
