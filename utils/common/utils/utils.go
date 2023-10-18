package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func PreSignedData(appid string, ts int64, nonce string, data []byte) []byte {
	stringBuffer := bytes.Buffer{}
	dataToSignList := []string{appid, strconv.FormatInt(ts, 10), nonce}
	dataToSign := strings.Join(dataToSignList, ";")
	stringBuffer.WriteString(dataToSign)
	stringBuffer.WriteString(";")
	stringBuffer.Write(data)
	return stringBuffer.Bytes()
}

func Sign(key, data []byte) string {
	h := hmac.New(sha256.New, key)

	// Write Data to it
	h.Write(data)

	// Get result and encode as hexadecimal string
	digest := hex.EncodeToString(h.Sum(nil))

	return digest
}
