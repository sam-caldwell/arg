package arg

import (
	"flag"
)

// NewFlag - Create a new boolean flag argument
func NewFlag(name string, usage string) (o *Flag, err error) {

	if err = validName(&name); err != nil {
		return nil, err
	}

	return &(Flag{
		value: flag.Bool(name, true, usage),
	}), err

}
