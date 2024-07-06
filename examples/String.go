package main

import (
	"fmt"
	"github.com/sam-caldwell/arg"
	"os"
)

func main() {
	value, err := arg.String("value", "5", "value 0-10 default 5", ".+")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	arg.Parse()
	fmt.Printf("value:%s", *value)
}
