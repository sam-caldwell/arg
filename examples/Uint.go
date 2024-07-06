package main

import (
	"fmt"
	"github.com/sam-caldwell/arg"
	"os"
)

func main() {
	value, err := arg.Uint("value", 5, 3, 10, "value 0-10 default 5")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	arg.Parse()
	fmt.Printf("value:%v\n", *value)
}
