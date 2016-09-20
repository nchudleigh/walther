package walther

import (
	"bytes"
	"math"
	"math/rand"
	"time"
)

// BASE62 should always be 0-9,a-z,A-Z
const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const base = len(alphabet)

const (
	bits   = 6           // 6 bits to represent 'alphabet' (62 == 111110)
	mask   = 1<<bits - 1 // All 1-bits, as many as 'bits'  (111111)
	bitmax = 63 / bits   // # of bits we can take from rand.Int63()
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

// Generate3 a random Base62 length string
func Generate3(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for 'bitmax' letters!
	for i, cache, remain := n-1, src.Int63(), bitmax; i >= 0; {
		if remain == 0 {
			remain, cache = bitmax, src.Int63()
		}
		idx := int(cache & mask)
		if idx < base {
			b[i] = balphabet[idx]
			i--
		}
		cache >>= bits
		remain--
	}

	return string(b)
}

// Generate2 a random Base62 length string
func Generate2(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		// mask the random integer with the bit mask
		// 1010101001010101001010101010101001000011 MASK
		//                 101010101010101010100001
		// 0000000000000000001010101010101000000001 RESULT
		idx := int(rand.Int63() & mask)
		if idx < base {
			b[i] = balphabet[idx]
			i++
		}
	}
	return string(b)
}

// Generate a random Base62 length string
func Generate(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = balphabet[rand.Int63()%int64(base)]
	}
	return string(result)
}
