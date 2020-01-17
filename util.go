package jwt

import (
	"encoding/hex"
	"math/rand"
	"time"
)

// RandByte is used to get a list rand bytes as secret
func RandByte() []byte {
	rand.Seed(time.Now().Unix())
	var letters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	length := 25 + rand.Intn(8)
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
		for i > 0 && b[i] == b[i-1] {
			b[i] = letters[rand.Intn(len(letters))]
		}
	}
	return b
}

// UUID v4
func UUID() string {
	version := byte(4)
	uuid := make([]byte, 16)
	rand.Read(uuid)

	// Set version
	uuid[6] = (uuid[6] & 0x0f) | (version << 4)

	// Set variant
	uuid[8] = (uuid[8] & 0xbf) | 0x80

	buf := make([]byte, 36)
	var dash byte = '-'
	hex.Encode(buf[0:8], uuid[0:4])
	buf[8] = dash
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = dash
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = dash
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = dash
	hex.Encode(buf[24:], uuid[10:])

	return string(buf)
}
