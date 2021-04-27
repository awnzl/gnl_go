package gnl

import (
	"context"
	"io"
	"testing"
)

func TestGetNextLine(t *testing.T) {
	wants := []string{"first string", "second string", "third string"}

	ctx, cancel := context.WithCancel(context.Background())



	got, err := io.ReadAll(res.Body)

}
