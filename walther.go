package walther

import (
	"bytes"
	"math"
	"math/rand"
	"time"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base     = len(alphabet)
	bits     = 6           // 6 bits to represent 'alphabet' (62 == 111110)
	mask     = 1<<bits - 1 // All 1-bits, as many as 'bits'  (111111)
	bitmax   = 63 / bits   // # of bits we can take from rand.Int63()
)

// byte array of alphabet
var balphabet = []byte(alphabet)

// Encode a positive integer to a Base62 token
func Encode(num int) []byte {
	if num == 0 {
		return []byte{alphabet[0]}
	}
	// result byte arr
	var res []byte
	// remainder of modulo
	var rem int

	// until num is == 0
	for num != 0 {
		// calculate remainder
		rem = num % base
		// calculate quotient
		num = num / base
		// push alphabet[rem] into res
		res = append(res, alphabet[rem])
	}
	// reverse res
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

// Decode a Base62 token to its original positive integer
func Decode(token []byte) int {
	// calculate base (should be 62)
	tokenlen := len(token)
	num := 0
	idx := 0
	// until num is == 0
	for _, c := range token {
		// calculate remainder
		power := (tokenlen - (idx + 1))
		// calculate quotient
		index := bytes.IndexByte(balphabet, c)
		// sum num and decode algo
		num += index * int(math.Pow(float64(base), float64(power)))
		// increment index token
		idx++
	}

	return num
}

var src = rand.NewSource(time.Now().UnixNano())

// Generate a random Base62 length string
func Generate(length int) string {
	result := make([]byte, length)
	for i, cache, remain := length-1, src.Int63(), bitmax; i >= 0; {
		if remain == 0 {
			remain, cache = bitmax, src.Int63()
		}
		idx := int(cache & mask)
		if idx < base {
			result[i] = balphabet[idx]
			i--
		}
		cache >>= bits
		remain--
	}

	return string(result)
}
