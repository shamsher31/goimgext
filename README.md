# goimgext

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/goimgext)
[![Go Build](https://github.com/shamsher31/goimgext/actions/workflows/go.yml/badge.svg)](https://github.com/shamsher31/goimgext/actions/workflows/go.yml)

List of Image file extensions for Go

### How to install
```go
go get github.com/shamsher31/goimgext
```

### How to use
```go
package main

import (
	"fmt"
	"github.com/shamsher31/goimgext"
)

func main() {
	fmt.Println(imgext.Get())
}
```

### Why
This package is inspired by [image-extensions](https://www.npmjs.com/package/image-extensions) npm module.

# License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
