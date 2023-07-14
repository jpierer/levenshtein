# levenshtein

An implementation of the Levenshtein algorithm in Go.

## Install

    go get github.com/jpierer/levenshtein

## Example

```go
package main

import (
	"fmt"

	"github.com/jpierer/levenshtein"
)

func main() {
	distance := levenshtein.Distance("foobar", "barbaz")
	fmt.Println(distance) // Output: 4
}

```

### Support Me

Give a ⭐ if this project was helpful in any way!

### License

The code is released under the [MIT LICENSE](/LICENSE).
