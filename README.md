# Go Constants `cnst`

[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/cnst)](https://goreportcard.com/report/github.com/speedyhoon/cnst)

Go string constants for [HTTP Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers)
and [Common MIME Types](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types)

## Install:

```sh
go get github.com/speedyhoon/cnst
```

## Examples:

```go
package main

import "github.com/speedyhoon/cnst/hdr"

func main() {
	println(hdr.AcceptEncoding, hdr.ContentType, hdr.CSP)
	// "accept-encoding" "content-type" "content-security-policy"
}
```

MIME Types are for use in the HTTP Content-Type header and contain a UTF-8 charset where appropriate.

```go
package main

import "github.com/speedyhoon/cnst/mime"

func main() {
	println(mime.CSS, mime.HTML, mime.JS, mime.WEBMv, mime.WEBP, mime.WOFF2)
	// "text/css; charset=utf8" "text/html; charset=utf8" "text/javascript" "video/webm" "image/webp" "font/woff2"
}
```

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/speedyhoon/cnst"
	"github.com/speedyhoon/cnst/hdr"
	"github.com/speedyhoon/cnst/mime"
	"github.com/speedyhoon/utl"
)

func main() {
	dirCSS := filepath.Join(utl.Cwd(), "css")
	cacheControl := fmt.Sprintf("public, max-age=%d", int((time.Hour * 24 * 365).Seconds()))

	http.HandleFunc("/css/",
		func(w http.ResponseWriter, r *http.Request) {
			// CSS content with UTF-8 encoding.
			w.Header().Set(hdr.ContentType, mime.CSS)

			// Cache expires in 1 year.
			w.Header().Set(hdr.CacheControl, cacheControl)

			// Serve file with Brotli compression.
			w.Header().Set(hdr.ContentEncoding, cnst.Brotli)
			http.FileServer(http.Dir(dirCSS)).ServeHTTP(w, r)
		})

	err := http.ListenAndServe(":", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
```

# Go standard library constants

All HTTP Methods (`GET`, `POST`, etc.) are defined in [`net/http`](https://golang.org/src/net/http/method.go).

All HTTP Status codes are defined in [`net/http`](https://golang.org/src/net/http/status.go).

Dynamic CPU Word size constants are defined in packages [`strconv`](https://golang.org/src/strconv/atoi.go)
and [`math/bits`](https://golang.org/src/math/bits/bits.go)

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
CPU Word size: 64 bits or 64 bits
value: 418, name: I'm a teapot, type: int
```
