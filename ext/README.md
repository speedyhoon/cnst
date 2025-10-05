# ext

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/cnst/ext.svg)](https://pkg.go.dev/github.com/speedyhoon/cnst/ext)
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/cnst/ext)](https://goreportcard.com/report/github.com/speedyhoon/cnst/ext)

Go string constants for commonly used file extensions.

```go
package main

import "github.com/speedyhoon/cnst/ext"

func main() {
	print("favicon" + ext.ICO)
}
```
Prints:<br>
`favicon.ico`
