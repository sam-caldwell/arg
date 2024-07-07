package main

import (
	"fmt"
	"github.com/sam-caldwell/arg"
	"os"
)

func main() {

	const name = "debug"

	debug, err := arg.NewFlag(name, "display debug messages")
	if err != nil {
		fmt.Printf("flag[%s]:%v", name, err)
		os.Exit(1)
	}

	arg.Parse()

	if err = debug.Verify(); err != nil {
		fmt.Printf("verify should return nil")
		os.Exit(1)
	}

	fmt.Printf("debug:%v\n", *debug.Value())
}
