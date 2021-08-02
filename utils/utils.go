package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"math"
	"math/rand"
	"time"
)

// Just some seed for random
func init() {
	rand.Seed(time.Now().UnixNano())
}

// NumberWLen returns number with desired length
func NumberWLen(length int) int {
	min := int(math.Pow10(length - 1))
	max := int(math.Pow10(length)) - 1
	return rand.Intn(max-min) + min
}

// HashString creates hash from string
func HashString(income string) string {
	return Hash([]byte(income))
}

// HashInt64 creates hash from int64
func HashInt64(income int64) string {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(income))
	return Hash(b)
}

// Hash creates hash from bytes :)
func Hash(income []byte) string {
	hasher := sha1.New()
	hasher.Write(income)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

// BtoI converts boolean to int: true -> 1, false -> 0
func BtoI(income bool) int {
	if income {
		return 1
	} else {
		return 0
	}
}

// SumBtoI finds sum of converted booleans
func SumBtoI(income ...bool) int {
	var res int
	for _, x := range income {
		res += BtoI(x)
	}
	return res
}
