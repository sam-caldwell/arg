package arg

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/errors"
	"github.com/sam-caldwell/file"
)

type ExistsOrNot bool

const (
	Exists    ExistsOrNot = true
	NotExists ExistsOrNot = false
)

// Filename - get a filename from the command-line and verify if it exists, if requireExists is true
func Filename(name, defaultValue string, requireExists ExistsOrNot, usage string) (*string, error) {
	value := flag.String(name, defaultValue, usage)
	if bool(requireExists) && !file.Exists(*value) {
		return nil, fmt.Errorf(errors.NotFound+errors.Details, name+":"+*value)
	}
	return value, nil
}
