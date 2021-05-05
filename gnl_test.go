package gnl

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadGetNextLine(t *testing.T) {
	want := [][]byte{
		[]byte("first string"),
		[]byte("second string"),
		[]byte("third string"),
	}

	ctx := context.Background()
	r := strings.NewReader("first string\nsecond string\nthird string")
	ch := GetNextLine(ctx, r)

	var got [][]byte
	for v := range ch {
		got = append(got, v)
	}

	for i := 0; i < len(want); i++ {
		assert.Equal(t, want[i], got[i])
	}
}

func TestCancelGetNextLine(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	r := strings.NewReader("first string\nsecond string\nthird string")
	ch := GetNextLine(ctx, r)

	assert.Equal(t, []byte("first string"), <-ch)

	cancel()
	v, isOpen := <-ch
	assert.Equal(t, []byte(nil), v)
	assert.Equal(t, false, isOpen)
	assert.Equal(t, context.Canceled, ctx.Err())
}

func TestReadGetNextLineErr(t *testing.T) {
	want := [][]byte{
		[]byte("first string"),
		[]byte("second string"),
		[]byte("third string"),
	}

	ctx := context.Background()
	r := strings.NewReader("first string\nsecond string\nthird string")
	ch, errCh := GetNextLineErr(ctx, r)

	var got [][]byte
	Loop:
	for {
		select {
		case err, ok := <-errCh:
			if ok {
				assert.Fail(t, fmt.Sprintf("unexpected error: %v", err))
				return
			}
		case bts, ok := <-ch:
			if !ok {
				break Loop
			}
			got = append(got, bts)
		}
	}

	for i := 0; i < len(want); i++ {
		assert.Equal(t, want[i], got[i])
	}
}

func TestCancelGetNextLineErr(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	r := strings.NewReader("first string\nsecond string\nthird string")
	ch, errCh := GetNextLineErr(ctx, r)

	assert.Equal(t, []byte("first string"), <-ch)

	cancel()
	want := context.Canceled
	got := <-errCh
	assert.Equal(t, want, got)

	v, isOpen := <-ch
	assert.Equal(t, []byte(nil), v)
	assert.Equal(t, false, isOpen)
}
