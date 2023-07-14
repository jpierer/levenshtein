# levenshtein
An implementation of the Levenshtein algorithm in Go.

Install
-------

    go get github.com/julianpierer/levenshtein

Example
-------

```go
package main

import (
	"fmt"

	"github.com/julianpierer/levenshtein"
)

func main() {
	distance := levenshtein.Distance("foobar", "barbaz")
	fmt.Println(distance) // Output: 4
}

```
