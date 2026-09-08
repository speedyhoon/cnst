# hdr

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/cnst/hdr.svg)](https://pkg.go.dev/github.com/speedyhoon/cnst/hdr)
[![Go Report Card](https://raw.githubusercontent.com/speedyhoon/speedyhoon/refs/heads/main/goReport.svg)](https://goreportcard.com/report/github.com/speedyhoon/cnst/hdr)
![license MIT](https://raw.githubusercontent.com/speedyhoon/speedyhoon/refs/heads/main/MIT.svg)

Go string constants for [HTTP Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers).

```go
package main

import "github.com/speedyhoon/cnst/hdr"

func main() {
	println(hdr.CSP)
	println(hdr.CrossOriginResourcePolicy)
}
```

Prints:

```
content-security-policy
cross-origin-resource-policy
```
