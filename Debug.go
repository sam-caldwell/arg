package arg

import "flag"

// Debug - Return true if -debug flag is passed to command-line
func Debug() *bool {
	value := flag.Bool("debug", false, "set the debug flag")
	return value
}
