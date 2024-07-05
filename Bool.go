package arg

import (
	"flag"
)

// Bool - get bool from the command-line and validate it against the given parameters
func Bool(name string, defaultValue bool, usage string) (*bool, error) {
	value := flag.Bool(name, defaultValue, usage)
	return value, nil
}
