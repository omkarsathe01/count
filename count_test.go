package count_test

import (
	"bytes"
	"github.com/rogpeppe/go-internal/testscript"
	"testing"

	"github.com/omkarsathe01/count"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"count": count.Main,
	})
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestLinesCountsLinesInInput(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("Line 1\nLine 2\nLine 3\nLine 4\n")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 4
	got := c.Lines()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestWithInputFromArgs_SetsInputToGivenPath(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.WithInputArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgs_IgnoresEmptyArgs(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3\n")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
		count.WithInputArgs([]string{}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}
