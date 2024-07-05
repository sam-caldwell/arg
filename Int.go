package arg

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Int - get int from the command-line and validate it against the given parameters
func Int(name string, defaultValue, min, max int, usage string) (*int, error) {
	value := flag.Int(name, defaultValue, usage)
	if *value <= min || *value >= max {
		return nil, fmt.Errorf(errors.InvalidInput+errors.Details, name)
	}
	return value, nil
}
