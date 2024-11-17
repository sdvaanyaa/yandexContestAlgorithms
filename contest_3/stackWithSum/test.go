package main

import (
	"fmt"
	"os"
)

func main() {
	out, err := os.Create("test_input.txt")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	n := 100000
	_, _ = fmt.Fprintln(out, n)

	for i := 0; i < n-1; i++ {
		_, _ = fmt.Fprintf(out, "+%d\n", 1000000000)
	}
	_, _ = fmt.Fprintf(out, "?%d\n", n-1)
}
