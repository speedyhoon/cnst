# hdr

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/cnst/hdr.svg)](https://pkg.go.dev/github.com/speedyhoon/cnst/hdr)
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/cnst/hdr)](https://goreportcard.com/report/github.com/speedyhoon/cnst/hdr)

Go string constants for [HTTP Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers).

```go
package main

import "github.com/speedyhoon/cnst/hdr"

func main() {
	println(hdr.CrossOriginResourcePolicy)
}
```
Prints:<br>
`cross-origin-resource-policy`
