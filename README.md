# file2go


```txt
┌─┐┬ ┬  ┌─┐   ┌─┐┌─┐
├┤ │ │  ├┤    │ ┬│ │
└  ┴ ┴─┘└─┘ 2 └─┘└─┘
https://github.com/lucasepe/file2go

Convert any file to Go source.

SYNOPSIS:

  file2go [string] [string]

DESCRIPTION:

The file2go utility reads a file from stdin and writes it to stdout,
converting each byte to its hex representation on the fly.

  * if the first [string] is present, it is printed before the data
  * if the second [string] is present, it is printed after the data

This program is used to embed binary or other files into Go source
files, for instance as a []byte.

EXAMPLES:

  date | file2go 'var myDate = []byte {' '}'

will produce:

  var myDate = []byte {
      0x46, 0x72, 0x69, 0x20, 0x4d, 0x61, 0x79, 0x20, 0x33, 0x31,
      0x20, 0x31, 0x37, 0x3a, 0x31, 0x38, 0x3a, 0x34, 0x38, 0x20,
      0x43, 0x45, 0x53, 0x54, 0x20, 0x32, 0x30, 0x32, 0x34, 0x0a
  }
```

<br/><br/><br/><br/>
---

I know very well about the go [package embed](https://pkg.go.dev/embed), but I wanted something _'old school'_.

If your workflow involves embedding multiple files, regenerating frequently your embedded data, then you definitely have to choose the [embed](https://pkg.go.dev/embed) package.
