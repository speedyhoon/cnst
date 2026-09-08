# edg

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/cnst/edg.svg)](https://pkg.go.dev/github.com/speedyhoon/cnst/edg)
[![Go Report Card](https://raw.githubusercontent.com/speedyhoon/speedyhoon/refs/heads/main/goReport.svg)](https://goreportcard.com/report/github.com/speedyhoon/cnst/edg)
![license MIT](https://raw.githubusercontent.com/speedyhoon/speedyhoon/refs/heads/main/MIT.svg)

Go integer constants for [polygons with specific quantities of edges](https://en.wikipedia.org/wiki/List_of_two-dimensional_geometric_shapes#Polygons_with_specific_numbers_of_sides) (sides, vertices, points or corners).

```go
package main

import "github.com/speedyhoon/cnst/edg"

func main() {
	println(edg.Octagon)
	println(edg.Heptagram)
}
```

Prints:

```
8
14
```
