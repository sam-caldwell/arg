package main

import (
	"fmt"
	"github.com/sam-caldwell/arg"
	"os"
)

func main() {
	value, err := arg.Choices("value", "a", "choices a,b,c,d", "a", "b", "c", "d")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	arg.Parse()
	fmt.Printf("value:%v\n", *value)
}
