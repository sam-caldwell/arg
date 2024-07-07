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

	file1, err := arg.Filename("invalid", invalidFile, arg.Exists, "file_usage")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	arg.Parse()

	fmt.Printf("value:%s", *file1)

}
