package file2go

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Options struct {
	Prefix string
	Suffix string
	In     io.Reader
	Out    io.Writer
	Indent int
}

func Do(opts Options) error {
	const (
		cols = 10
	)

	var (
		indent string
		m      int64
		b      byte
		err    error
		buf    *bufio.Reader //bytes.Buffer
	)

	if opts.Indent > 0 {
		indent = strings.Repeat(" ", opts.Indent)
	}

	if opts.In == nil {
		opts.In = os.Stdin
	}

	if opts.Out == nil {
		opts.Out = os.Stdout
	}

	if len(opts.Prefix) > 0 {
		fmt.Fprintln(opts.Out, opts.Prefix)
		fmt.Fprint(opts.Out, indent)
	}

	buf = bufio.NewReader(opts.In)
	for {
		b, err = buf.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		doIndent := false
		_, err = buf.Peek(1)
		if err == nil {
			fmt.Fprintf(opts.Out, "0x%02x, ", b)
			doIndent = true
		} else {
			fmt.Fprintf(opts.Out, "0x%02x", b)
		}

		m++
		if m%cols == 0 {
			fmt.Fprintln(opts.Out)
			if doIndent {
				fmt.Fprint(opts.Out, indent)
			}
		}
	}

	if len(opts.Suffix) > 0 {
		if m%cols > 0 {
			fmt.Fprintln(opts.Out)
		}
		fmt.Fprintln(opts.Out, opts.Suffix)
	}

	return nil
}
