# cnst
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/cnst)](https://goreportcard.com/report/github.com/speedyhoon/cnst)

Go string constants for [HTTP Methods](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods), [HTTP Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers) and [Common MIME Types](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types)
```go
cnst.Get  = "GET"
cnst.Post = "POST"
cnst.Put  = "PUT"
```

```go
cnst.AcceptEncoding = "accept-encoding"
cnst.ContentType    = "content-type"
```

MIME Types are for use in the HTTP Content-Type header and contain a UTF-8 charset where appropriate.
```go
cnst.MimeCSS   = "text/css; charset=utf8"
cnst.MimeHTM   = "text/html; charset=utf8"
cnst.MimeMJS   = "text/javascript"
cnst.MimeWEBM  = "video/webm"
cnst.MimeWEBP  = "image/webp"
cnst.MimeWOFF2 = "font/woff2"
```

## Example Usage:
```sh
go get github.com/speedyhoon/cnst
```
```go
package main

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/speedyhoon/cnst"
	"github.com/speedyhoon/utl"
)

func main() {
	dirCSS := filepath.Join(utl.Cwd(), "css")

	http.HandleFunc("/css/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(cnst.ContentType, cnst.MimeCSS)

			//Cache headers:
			w.Header().Set(cnst.CacheControl, "public")
			//Cache expires in 1 year
			w.Header().Set(cnst.Expires, time.Now().UTC().AddDate(1, 0, 0).Format("Mon, 02 Jan 2006 15:04:05 GMT"))
			w.Header().Set(cnst.Vary, cnst.AcceptEncoding)

			//Serve file with Brotli compression
			w.Header().Set(cnst.ContentEncode, cnst.Brotli)
			http.FileServer(http.Dir(dirCSS)).ServeHTTP(w, r)
		})
	http.ListenAndServe(":", nil)
}
```
\
\
\
HTTP Status codes 100 - 511 are already constants defined in [package `net/http`](https://golang.org/src/net/http/status.go)
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("value: %d, name: %s, type: %T", http.StatusTeapot, http.StatusText(http.StatusTeapot), http.StatusTeapot)
}

//Outputs: value: 418, name: I'm a teapot, type: int
```
