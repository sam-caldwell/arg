package main

import (
	"fmt"
	"github.com/sam-caldwell/arg"
	"os"
)

func main() {
	value, err := arg.NewInt("value", 0, -5, 10, "value -5...10 default 0")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	arg.Parse()

	if err = value.Verify(); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("value:%v\n", *value.Value())
}
