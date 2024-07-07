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

	fmt.Printf("value:%v\n", *value)
}

//
//func main() {
//
//	const name = "debug"
//
//	debug, err := arg.NewFlag(name, "display debug messages")
//	if err != nil {
//		fmt.Printf("flag[%s]:%v", name, err)
//	}
//
//	arg.Parse()
//
//	if err = debug.Verify(); err != nil {
//		fmt.Printf("verify should return nil")
//	}
//
//	fmt.Printf("debug:%v\n", *debug.Value())
//}
