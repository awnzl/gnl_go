package gnl

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNextLine(t *testing.T) {
	wants := [][]byte{
		[]byte("first string"),
		[]byte("second string"),
		[]byte("third string"),
	}

	ctx, cancel := context.WithCancel(context.Background())
	r := strings.NewReader("first string\nsecond string\nthird string")
	ch := GetNextLine(ctx, r)

	for i := 0; i < len(wants); i++ {
		assert.Equal(t, wants[i], <-ch)
	}

	cancel()
	canceled := GetNextLine(ctx, strings.NewReader("first string\nsecond string"))
	_, isOpened := <-canceled
	assert.Equal(t, false, isOpened)
	assert.Equal(t, errors.New("context canceled"), ctx.Err())
}

func TestGetNextLineErr(t *testing.T) {
	wants := [][]byte{
		[]byte("first string"),
		[]byte("second string"),
		[]byte("third string"),
	}

	ctx, cancel := context.WithCancel(context.Background())
	r := strings.NewReader("first string\nsecond string\nthird string")
	ch, errc := GetNextLineErr(ctx, r)

	for i := 0; i < len(wants); i++ {
		assert.Equal(t, wants[i], <-ch)
	}
	assert.Equal(t, nil, <-errc)

	cancel()
	canceled, errc := GetNextLineErr(ctx, strings.NewReader("first string"))
	_, isOpened := <-canceled
	assert.Equal(t, false, isOpened)
	assert.Equal(t, errors.New("context canceled"), <-errc)
}
