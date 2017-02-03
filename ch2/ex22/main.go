package main

import (
	"bufio"
	"fmt"
	"rgogunskiy/golangbook/ch2/ex22/tempconv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(0)
			}

			convertAndPrint(t)

		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		var text string
		for true {
			fmt.Print("Enter a nubmer: ")
			scanner.Scan()
			text = scanner.Text()
			if text == "q" {
				break
			}
			number, err := strconv.ParseFloat(text, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s is not a number\n", text)
				continue
			}
			convertAndPrint(number)
		}
	}
}

func convertAndPrint(t float64) {
	{
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Println("== Fagrenheit <-> Celsius ==")
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
	{
		f := tempconv.Feet(t)
		m := tempconv.Meters(t)
		fmt.Println("== Feet <-> Meters ==")
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToM(f), m, tempconv.MToF(m))
	}
	{
		p := tempconv.Pounds(t)
		k := tempconv.Kilograms(t)
		fmt.Println("== Pounds <-> Kilograms ==")
		fmt.Printf("%s = %s, %s = %s\n",
			p, tempconv.PToK(p), k, tempconv.KToP(k))
	}
}
