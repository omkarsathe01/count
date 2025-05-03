package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
	input  io.Reader
	output io.Writer
}

type option func(*counter) error

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithInputArgs(s []string) option {
	return func(c *counter) error {
		if len(s) < 1 {
			return nil
		}
		file, err := os.Open(s[0])
		if err != nil {
			return err
		}
		c.input = file
		return nil
	}
}

func WithOutput(w io.Writer) option {
	return func(c *counter) error {
		if w == nil {
			return errors.New("nil output writer")
		}
		c.output = w
		return nil
	}
}

func NewCounter(ops ...option) (*counter, error) {
	c := &counter{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, op := range ops {
		err := op(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *counter) Lines() int {
	lines := 0
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		lines++
	}
	return lines
}

func Main() {
	c, err := NewCounter(
		WithInputArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(c.Lines())
}
