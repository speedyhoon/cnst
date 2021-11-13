# Go Constants `cnst`
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/cnst)](https://goreportcard.com/report/github.com/speedyhoon/cnst)

Go string constants for [HTTP Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers) and [Common MIME Types](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types)

```go
package main

import (
	"fmt"
	
	"github.com/speedyhoon/cnst/hdrs"
)

func main() {
	fmt.Println(hdrs.AcceptEncoding, hdrs.ContentType)
	// "accept-encoding" "content-type"
}
```

MIME Types are for use in the HTTP Content-Type header and contain a UTF-8 charset where appropriate.
```go
cnst.MimeCSS   = "text/css; charset=utf8"
cnst.MimeHTM   = "text/html; charset=utf8"
cnst.MimeJS    = "text/javascript"
cnst.MimeWEBM  = "video/webm"
cnst.MimeWEBP  = "image/webp"
cnst.MimeWOFF2 = "font/woff2"
```

## Install:
```sh
go get github.com/speedyhoon/cnst
```

## Example Usage:
```go
package main

import (
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/speedyhoon/cnst"
	"github.com/speedyhoon/cnst/hdrs"
	"github.com/speedyhoon/utl"
)

func main() {
	dirCSS := filepath.Join(utl.Cwd(), "css")

	http.HandleFunc("/css/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(hdrs.ContentType, cnst.MimeCSS)

			//Cache headers:
			w.Header().Set(hdrs.CacheControl, "public")
			//Cache expires in 1 year
			w.Header().Set(hdrs.Expires, time.Now().UTC().AddDate(1, 0, 0).Format("Mon, 02 Jan 2006 15:04:05 GMT"))
			w.Header().Set(hdrs.Vary, hdrs.AcceptEncoding)

			//Serve file with Brotli compression
			w.Header().Set(hdrs.ContentEncoding, cnst.Brotli)
			http.FileServer(http.Dir(dirCSS)).ServeHTTP(w, r)
		})

	err := http.ListenAndServe(":", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
```

\
\
\
HTTP Methods (`GET`, `POST`, etc.) are already constants defined
in [package `net/http`](https://golang.org/src/net/http/status.go)

HTTP Status codes 100 - 511 are already constants defined
in [package `net/http`](https://golang.org/src/net/http/method.go)
```go
package main

import (
	"fmt"
	"math/bits"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println(
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	)
	fmt.Println("CPU Word size:", strconv.IntSize, "bits or", bits.UintSize, "bits")
	fmt.Printf("value: %d, name: %s, type: %[1]T\n", http.StatusTeapot, http.StatusText(http.StatusTeapot))
}
```
Outputs:
```
GET HEAD POST PUT PATCH DELETE CONNECT OPTIONS TRACE
value: 418, name: I'm a teapot, type: int
```
