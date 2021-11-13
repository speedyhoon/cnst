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
package main

import (
	"fmt"

	"github.com/speedyhoon/cnst/mime"
)

func main() {
	fmt.Println(mime.CSS, mime.HTML, mime.JS, mime.WEBMv, mime.WEBP, mime.WOFF2)
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
	"github.com/speedyhoon/cnst/hdrs"
	"github.com/speedyhoon/cnst/mime"
	"github.com/speedyhoon/utl"
)

func main() {
	dirCSS := filepath.Join(utl.Cwd(), "css")
	cacheControl := fmt.Sprintf("public, max-age=%d", int((time.Hour * 24 * 365).Seconds()))

	http.HandleFunc("/css/",
		func(w http.ResponseWriter, r *http.Request) {
			// CSS content with UTF-8 encoding.
			w.Header().Set(hdrs.ContentType, mime.CSS)

			// Cache expires in 1 year.
			w.Header().Set(hdrs.CacheControl, cacheControl)

			// Serve file with Brotli compression.
			w.Header().Set(hdrs.ContentEncoding, cnst.Brotli)
			http.FileServer(http.Dir(dirCSS)).ServeHTTP(w, r)
		})

	err := http.ListenAndServe(":", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
```

# Go standard library constants

All HTTP Methods (`GET`, `POST`, etc.) are defined in [`net/http`](https://golang.org/src/net/http/status.go).

All HTTP Status codes are defined in [`net/http`](https://golang.org/src/net/http/method.go).

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
	fmt.Printf("value: %d, name: %s, type: %[1]T\n", http.StatusTeapot, http.StatusText(http.StatusTeapot))
	fmt.Println("CPU Word size:", strconv.IntSize, "bits or", bits.UintSize, "bits")
}
```
Outputs:
```
GET HEAD POST PUT PATCH DELETE CONNECT OPTIONS TRACE
value: 418, name: I'm a teapot, type: int
CPU Word size: 64 bits or 64 bits 
```
