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

		scanner := bufio.NewScanner(r)

		for {
			if !scanner.Scan() {
				break
			}

			select {
			case <-ctx.Done():
				return
			case out <- scanner.Bytes():
			}
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

		scanner := bufio.NewScanner(r)

		for {
			if !scanner.Scan() {
				break
			}

			if err := scanner.Err(); err != nil {
				errc <- err
				break
			}

			select {
			case <-ctx.Done():
				errc <-ctx.Err()
				return
			case out <-scanner.Bytes():
			}
		}
	}()

	return out, errc
}
