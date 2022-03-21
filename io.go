package kyopro

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type IO struct {
	sc *bufio.Scanner
	w  *bufio.Writer
}

func NewIO() *IO {
	return NewIOFrom(os.Stdin, os.Stdout)
}

func NewIOFrom(src io.Reader, dst io.Writer) *IO {
	sc := bufio.NewScanner(src)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(dst)
	return &IO{sc, w}
}

// めっちゃ遅いです
func Next[T any](io *IO) T {
	var t T
	io.scan()
	fmt.Sscan(io.sc.Text(), &t)
	return t
}

func (io *IO) NextInt64() int64 {
	io.scan()
	val, err := strconv.ParseInt(io.sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return val
}

func (io *IO) NextInt() int {
	return int(io.NextInt64())
}

func (io *IO) scan() {
	if !io.sc.Scan() {
		panic("")
	}
}

func (io *IO) Flush() error {
	return io.w.Flush()
}

func (io *IO) Println(a ...interface{}) error {
	_, err := io.w.WriteString(fmt.Sprintln(a...))
	return err
}
