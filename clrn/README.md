# clrn

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/cnst/clrn.svg)](https://pkg.go.dev/github.com/speedyhoon/cnst/clrn)
[![Go Report Card](https://raw.githubusercontent.com/speedyhoon/speedyhoon/refs/heads/main/goReport.svg)](https://goreportcard.com/report/github.com/speedyhoon/cnst/clrn)
![license MIT](https://raw.githubusercontent.com/speedyhoon/speedyhoon/refs/heads/main/MIT.svg)

Go string constants for [CSS named-colors](https://developer.mozilla.org/en-US/docs/Web/CSS/named-color).

```go
package main

import "github.com/speedyhoon/cnst/clrn"

func main() {
	println(clrn.LightGoldenrodYellow)
	println(clrn.RebeccaPurple)
}
```

Prints:

```
lightgoldenrodyellow
rebeccapurple
```
