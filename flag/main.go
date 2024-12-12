package main

import (
	"flag"
	"fmt"
)

var intFlag int64
var stringFlag string
var boolFlag bool

func init() {
	flag.Int64Var(&intFlag, "int", 0, "int flag")
	flag.StringVar(&stringFlag, "string", "", "string flag")
	flag.BoolVar(&boolFlag, "bool", false, "bool flag")
}

func main() {
	flag.Parse()

	fmt.Printf("int: %d\n", intFlag)
	fmt.Printf("string: %s\n", stringFlag)
	fmt.Printf("bool: %t\n", boolFlag)
}
