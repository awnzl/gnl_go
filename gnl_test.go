package gnl

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadGetNextLine(t *testing.T) {
	wants := [][]byte{
		[]byte("first string"),
		[]byte("second string"),
		[]byte("third string"),
	}

	ctx, _ := context.WithCancel(context.Background())
	r := strings.NewReader("first string\nsecond string\nthird string")
	ch := GetNextLine(ctx, r)

	var got [][]byte
	for v := range ch {
		got = append(got, v)
	}

	for i := 0; i < len(wants); i++ {
		assert.Equal(t, wants[i], got[i])
	}
}

func TestCancelGetNextLine(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	r := strings.NewReader("first string\nsecond string\nthird string")
	ch := GetNextLine(ctx, r)

	assert.Equal(t, []byte("first string"), <-ch)

	cancel()
	v, isOpened := <-ch
	assert.Equal(t, []byte(nil), v)
	assert.Equal(t, false, isOpened)
	assert.Equal(t, errors.New("context canceled"), ctx.Err())
}

func TestReadGetNextLineErr(t *testing.T) {
	wants := [][]byte{
		[]byte("first string"),
		[]byte("second string"),
		[]byte("third string"),
	}

	ctx, _ := context.WithCancel(context.Background())
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

	for i := 0; i < len(wants); i++ {
		assert.Equal(t, wants[i], got[i])
	}
}

func TestCancelGetNextLineErr(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	r := strings.NewReader("first string\nsecond string\nthird string")
	ch, errCh := GetNextLineErr(ctx, r)

	assert.Equal(t, []byte("first string"), <-ch)

	cancel()
	want := errors.New("context canceled")
	got := <-errCh
	assert.Equal(t, want, got)

	v, isOpened := <-ch
	assert.Equal(t, []byte(nil), v)
	assert.Equal(t, false, isOpened)
}
