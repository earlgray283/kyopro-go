package kyopro

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"testing"
)

func makeInputCase(n int) io.Reader {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("%v\n", n))
	for i := 0; i < n; i++ {
		x, y := rand.Int63n(1<<63-1), rand.Int63n(1<<63-1)
		buf.WriteString(fmt.Sprintf("%v %v\n", x, y))
	}
	return bytes.NewReader(buf.Bytes())
}

func BenchmarkIOwithSscan(b *testing.B) {
	inputFile := makeInputCase(b.N % 1_000_001)

	b.ResetTimer()

	io := NewIOFrom(inputFile, os.Stdout)
	defer io.Flush()
	n := Next[int](io)
	sx := make([]int, n)
	sy := make([]int, n)
	for i := 0; i < n; i++ {
		x, y := Next[int](io), Next[int](io)
		sx[i] = x
		sy[i] = y
	}
}

func BenchmarkIOwithParseInt(b *testing.B) {
	inputFile := makeInputCase(b.N % 1_000_001)

	b.ResetTimer()

	io := NewIOFrom(inputFile, os.Stdout)
	defer io.Flush()
	n := Next[int](io)
	sx := make([]int, n)
	sy := make([]int, n)
	for i := 0; i < n; i++ {
		x, y := io.NextInt(), io.NextInt()
		sx[i] = x
		sy[i] = y
	}
}
