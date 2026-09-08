# mime

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/cnst/mime.svg)](https://pkg.go.dev/github.com/speedyhoon/cnst/mime)
[![Go Report Card](https://raw.githubusercontent.com/speedyhoon/speedyhoon/refs/heads/main/goReport.svg)](https://goreportcard.com/report/github.com/speedyhoon/cnst/mime)
![license MIT](https://raw.githubusercontent.com/speedyhoon/speedyhoon/refs/heads/main/MIT.svg)

Go string constants for [Common MIME Types](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types).

```go
package main

import "github.com/speedyhoon/cnst/mime"

func main() {
	println(mime.DOCX)
	println(mime.FLAC)
}
```

Prints:

```
application/vnd.openxmlformats-officedocument.wordprocessingml.document
audio/flac
```
