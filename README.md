# goimgext
List of Image file extensions for Go

# How to install
<pre>
go get github.com/shamsher31/goimgext
</pre>

# How to use
<pre>
package main

import (
	"fmt"
	"github.com/shamsher31/goimgext"
)

func main() {
	fmt.Println(imgext.Get())
}
</pre>

# Why
This package is inspired by [image-extensions](https://www.npmjs.com/package/image-extensions) npm module for image extension.

# License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
