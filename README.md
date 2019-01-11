# asPtr

[![GoDoc](https://godoc.org/github.com/zeeke/asptr?status.svg)](https://godoc.org/github.com/zeeke/asptr)

Get pointers from literals
(as explained by @icza [here](https://stackoverflow.com/a/30716481/1850410))

## Usage

```go

import "github.com/zeeke/asptr"

var a *int = asptr.Int(42)
var b *string = asptr.String("Hello World")
var c *bool = asptr.Bool(true)
var d *float32 = asptr.Float32(4.2)

```

