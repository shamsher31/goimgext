# goimgext

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/goimgext)
[![Build Status](https://travis-ci.org/shamsher31/goimgext.svg?branch=master)](https://travis-ci.org/shamsher31/goimgext)

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
