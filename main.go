package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	q := quote.Go()

	fmt.Println(q)
}
