# lang

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/cnst/lang.svg)](https://pkg.go.dev/github.com/speedyhoon/cnst/lang)
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/cnst/lang)](https://goreportcard.com/report/github.com/speedyhoon/cnst/lang)

Go string constants for most [HTTP languages](https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry) 
for use in HTTP headers [`Accept-Language`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Accept-Language),
[`Content-Language`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Content-Language) and
[HTML global attribute `lang`](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/lang).

```go
package main

import "github.com/speedyhoon/cnst/lang"

func main() {
	println(lang.Greenlandic)
	println(lang.WelshUnitedKingdom)
}
```
Prints:
```
kl
cy-GB
```
