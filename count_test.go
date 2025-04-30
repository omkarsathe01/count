package count_test

import (
	"bytes"
	"github.com/omkarsathe01/count"
	"testing"
)

func TestCountLines(t *testing.T) {
	t.Parallel()

	// write multiple lines to buffer
	buf := new(bytes.Buffer)
	buf.WriteString("Line 1\nLine 2\nLine 3\nLine 4\n")

	want := 4

	c, err := count.NewCounter(count.WithInput(buf))

	if err != nil {
		t.Fatal(err)
	}

	got := c.Lines()

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}
