package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lucasepe/file2go/internal/file2go"
)

const (
	banner = `┌─┐┬ ┬  ┌─┐   ┌─┐┌─┐
├┤ │ │  ├┤    │ ┬│ │
└  ┴ ┴─┘└─┘ 2 └─┘└─┘
https://github.com/lucasepe/file2go`
)

func main() {
	flag.Usage = usage
	flag.Parse()

	opts := file2go.Options{
		In:  os.Stdin,
		Out: os.Stdout,
	}

	switch flag.NArg() {
	case 1:
		opts.Prefix = flag.Arg(0)
		opts.Indent = 4
	case 2:
		opts.Prefix = flag.Arg(0)
		opts.Suffix = flag.Arg(1)
		opts.Indent = 4
	}

	err := file2go.Do(opts)
	if err != nil {
		exitOnErr(err)
	}
}

func usage() {
	fmt.Fprint(os.Stderr, banner)
	fmt.Fprint(os.Stderr, "\n\n")

	fmt.Fprintf(os.Stderr, "Convert any file to Go source.\n\n")

	fmt.Fprint(os.Stderr, "SYNOPSIS:\n\n")
	fmt.Fprintf(os.Stderr, "  %s [string] [string]\n\n", appName())

	fmt.Fprint(os.Stderr, "DESCRIPTION:\n\n")
	fmt.Fprintf(os.Stderr, "The %s utility reads a file from stdin and writes it to stdout,\n", appName())
	fmt.Fprint(os.Stderr, "converting each byte to its hex representation on the fly.\n\n")

	fmt.Fprint(os.Stderr, "  * if the first [string] is present, it is printed before the data\n")
	fmt.Fprint(os.Stderr, "  * if the second [string] is present, it is printed after the data\n\n")

	fmt.Fprint(os.Stderr, "This program is used to embed binary or other files into Go source\n")
	fmt.Fprint(os.Stderr, "files, for instance as a []byte.\n\n")

	fmt.Fprint(os.Stderr, "EXAMPLES:\n\n")
	fmt.Fprintf(os.Stderr, "  date | %s 'var myDate = []byte' '}'\n\n", appName())
	fmt.Fprintf(os.Stderr, "will produce:\n\n")
	fmt.Fprintf(os.Stderr, "  var myDate = []byte\n")
	fmt.Fprintf(os.Stderr, "      0x46, 0x72, 0x69, 0x20, 0x4d, 0x61, 0x79, 0x20, 0x33, 0x31,\n")
	fmt.Fprintf(os.Stderr, "      0x20, 0x31, 0x37, 0x3a, 0x31, 0x38, 0x3a, 0x34, 0x38, 0x20,\n")
	fmt.Fprintf(os.Stderr, "      0x43, 0x45, 0x53, 0x54, 0x20, 0x32, 0x30, 0x32, 0x34, 0x0a\n")
	fmt.Fprintf(os.Stderr, "  }\n")

	os.Exit(0)
}

func appName() string {
	return filepath.Base(os.Args[0])
}

// exitOnErr check for an error and eventually exit
func exitOnErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
