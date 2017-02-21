package main

import (
	"fmt"
	"os"
	"rgogunskiy/golangbook/ch2/ex24/popcount"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		x, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(0)
		}

		fmt.Println(popcount.PopCount(uint64(x)))
		fmt.Println(popcount.PopCount2(uint64(x)))
	}
}
