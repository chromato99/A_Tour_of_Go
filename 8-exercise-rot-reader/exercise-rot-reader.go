package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	bytesRead, err := r.r.Read(b)

	if err != nil {
		return bytesRead, io.EOF
	}

	for i := 0; i < bytesRead; i++ {
		switch {
		case 'A' <= b[i] && b[i] <= 'M':
			b[i] += 13
		case 'N' <= b[i] && b[i] <= 'Z':
			b[i] -= 13
		case 'a' <= b[i] && b[i] <= 'm':
			b[i] += 13
		case 'n' <= b[i] && b[i] <= 'z':
			b[i] -= 13
		}
	}

	return bytesRead, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
