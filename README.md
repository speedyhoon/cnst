# Go Constants `cnst`

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/cnst.svg)](https://pkg.go.dev/github.com/speedyhoon/cnst)
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/cnst)](https://goreportcard.com/report/github.com/speedyhoon/cnst)

Go constants for [HTTP Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers),
[HTTP languages](https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry),
[Common MIME Types](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types),
[W3C named-colors](https://developer.mozilla.org/en-US/docs/Web/CSS/named-color) and
how many [edges polygons have](https://en.wikipedia.org/wiki/Polygon#Naming).

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

			// Serve files with Brotli compression.
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
	"math"
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
	// Math constants.
	fmt.Println(math.E, math.Pi, math.Phi)
	fmt.Println(math.Sqrt2, math.SqrtE, math.SqrtPi, math.SqrtPhi)
	fmt.Println(math.Ln2, math.Log2E, math.Ln10, math.Log10E)
	fmt.Println(math.MaxFloat32, math.SmallestNonzeroFloat32)
	fmt.Println(math.MaxFloat64, math.SmallestNonzeroFloat64)
	fmt.Println(math.MaxInt, math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64)
	fmt.Println(math.MinInt, math.MinInt8, math.MinInt16, math.MinInt32, math.MinInt64)
	fmt.Println(uint64(math.MaxUint), math.MaxUint8, math.MaxUint16, math.MaxUint32, uint64(math.MaxUint64))
}
```
Outputs:
```
GET HEAD POST PUT PATCH DELETE CONNECT OPTIONS TRACE
CPU Word size: 64 bits or 64 bits
value: 418, name: I'm a teapot, type: int
2.718281828459045 3.141592653589793 1.618033988749895
1.4142135623730951 1.6487212707001282 1.772453850905516 1.272019649514069
0.6931471805599453 1.4426950408889634 2.302585092994046 0.4342944819032518
3.4028234663852886e+38 1.401298464324817e-45
1.7976931348623157e+308 5e-324
9223372036854775807 127 32767 2147483647 9223372036854775807
-9223372036854775808 -128 -32768 -2147483648 -9223372036854775808
18446744073709551615 255 65535 4294967295 18446744073709551615
```
