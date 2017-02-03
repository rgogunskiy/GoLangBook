package main

import (
	"fmt"
	"os"
	"rgogunskiy/golangbook/ch2/ex21/tempconv"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Println("== Fahrenheit <-> Celsius ==")
		fmt.Printf("%s = %s, %s = %s \n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))

		fmt.Println("== Kelvin <-> Celsius ==")
		fmt.Printf("%s = %s, %s = %s \n",
			k, tempconv.KToC(k), c, tempconv.CToK(c))
	}
}
