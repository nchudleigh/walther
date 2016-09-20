package walther

import (
	"bytes"
	"testing"
)

// Cases Shared by Encode and Decode
var cases = map[int][]byte{
	0:          []byte("0"),
	1:          []byte("1"),
	100:        []byte("1C"),
	999:        []byte("g7"),
	9999999999: []byte("aUKYOz"),
}

func TestEncode(t *testing.T) {
	for input, expected := range cases {
		output := Encode(input)
		if !bytes.Equal(expected, output) {
			t.Errorf("Encode(%d), Got: %d,  Expected: %d", input, output, expected)
		}
	}
}

func TestDecode(t *testing.T) {
	for expected, input := range cases {
		output := Decode(input)
		if expected != output {
			t.Errorf("Decode(%s) => Got: %d,  Expected: %d", input, output, expected)
		}
	}
}

func TestGenerate(t *testing.T) {
	var cases = []int{0, 1, 100, 999, 9999, 99999, 999999, 9999999, 99999999, 999999999}
	for _, length := range cases {
		output := Generate3(length)
		if len(output) != length {
			t.Errorf("Generate(%d) => Got: %d,  Expected: %d", length, len(output), length)
		}
	}
}

// a common length for the use case of generate
const n = 16

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate(n)
	}
}

func BenchmarkGenerate2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate2(n)
	}
}

func BenchmarkGenerate3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate3(n)
	}
}

func BenchmarkGenerate4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate4(n)
	}
}
