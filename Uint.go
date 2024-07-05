package arg

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Uint - get uint from the command-line and validate it against the given parameters
func Uint(name string, defaultValue, min, max uint, usage string) (*uint, error) {
	value := flag.Uint(name, defaultValue, usage)
	if *value <= min || *value >= max {
		return nil, fmt.Errorf(errors.InvalidInput+errors.Details, name)
	}
	return value, nil
}
