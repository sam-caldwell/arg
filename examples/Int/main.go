package main

import (
	"fmt"
	"github.com/sam-caldwell/arg"
	"os"
)

func main() {
	value, err := arg.Int("value", 0, -5, 10, "value -5-10 default 0")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	arg.Parse()
	fmt.Printf("value:%v\n", *value)
}
