package main

import (
	"fmt"
	"github.com/sam-caldwell/arg"
	"os"
)

func main() {

	const (
		invalidFile = "/tmp/invalid_file.txt"
	)

	file1, err := arg.NewFilename("file", arg.Exists, "file_usage")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	arg.Parse()

	if err = file1.Verify(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("value:%s", *file1.Value())

}
