package gnl

import (
	"bufio"
	"context"
	"io"
)

func GetNextLine(ctx context.Context, r io.Reader) <-chan []byte {
	out := make(chan []byte)

	go func() {
		defer close(out)

		select {
		case <-ctx.Done():
			return
		default:
			scan(r, out)
		}
	}()

	return out
}

func GetNextLineErr(ctx context.Context, r io.Reader) (<-chan []byte, <-chan error) {
	out := make(chan []byte)
	errc := make(chan error)

	go func() {
		defer close(out)
		defer close(errc)

		select {
		case <-ctx.Done():
			close(out)
			errc <-ctx.Err()
			return
		default:
			close(errc)
			scan(r, out)
		}
	}()

	return out, errc
}

func scan(r io.Reader, out chan<- []byte) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		out <- scanner.Bytes()
	}
}
