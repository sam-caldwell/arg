package arg

import (
	"flag"
)

// NewBool - Create a new Boolean argument
func NewBool(name string, defaultValue bool, usage string) (o *Bool, err error) {

	if err = validName(&name); err != nil {
		return nil, err
	}

	return &(Bool{
		value: flag.Bool(name, defaultValue, usage),
	}), err

}
