package arg

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Choices - Get a string from command-line and verify it is one of a set of choices
func Choices(name, defaultValue, usage string, choices ...string) (*string, error) {
	value := flag.String(name, defaultValue, usage)
	flag.Parse()
	for _, validValue := range choices {
		if *value == validValue {
			return value, nil
		}
	}
	return value, fmt.Errorf(errors.InvalidInput+errors.Details, *value)
}
