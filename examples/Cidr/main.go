package main

import (
	"fmt"
	"github.com/sam-caldwell/arg"
	"os"
)

func main() {
	value, err := arg.NewCidr("cidr", "0.0.0.0/0", "ipv4 cidr address")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	arg.Parse()

	if err = value.Verify(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s", *value.Value())
}
