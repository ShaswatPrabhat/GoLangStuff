package main

import (
	"code/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		argument, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("Error in parsing float input %v", err)
		}
		f := tempconv.Fahrenheit(argument)
		c := tempconv.Celsius(argument)
		fmt.Printf("%s = %s, %s=%s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
