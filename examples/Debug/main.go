package main

import (
	"fmt"
	"github.com/sam-caldwell/arg"
)

func main() {
	debug := arg.Debug()
	arg.Parse()
	fmt.Printf("debug:%v\n", *debug)
}
