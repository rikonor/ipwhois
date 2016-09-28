### ipwhois

Wrapper around [py-ipwhois](http://ipwhois.readthedocs.io/en/latest/).

To use this package you must have `Python` and `ipwhois` installed.

### Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/rikonor/ipwhois"
)

func main() {
	res, err := ipwhois.LookupIP("8.8.8.8")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", res)
}

```

### License

MIT
