package arg

import (
	"regexp"
)

// String - Represents a command-line string argument and optionally performs validation
type String struct {
	value      *string
	validation *regexp.Regexp
}
