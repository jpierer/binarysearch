# Binary search algorithm

An implementation of the binary search algorithm in Go

## Wikipedia:

[EN: Binary search algorithm](https://en.wikipedia.org/wiki/Binary_search_algorithm)

[DE: Binäre Suche](https://de.wikipedia.org/wiki/Bin%C3%A4re_Suche)

## Install

    go get github.com/jpierer/binarysearch

## Example

```go
package main

import (
	"fmt"
	"github.com/jpierer/binarysearch"
)

func main() {
	data := []int{44, 55, 66, 67, 68, 99, 100}
	indexPos, err := binarysearch.Search(data, 99)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(indexPos) // Output: 5
}

```

### Support Me

Give a ⭐ if this project was helpful in any way!

### License

The code is released under the [MIT LICENSE](/LICENSE).
