package arg

import (
	"flag"
)

// NewUint - Create a new 64-bit unsigned-integer flag argument
func NewUint(name string, defaultValue, minimum, maximum uint64, usage string) (o *Uint, err error) {

	if err = validName(&name); err != nil {
		return nil, err
	}

	return &(Uint{
		value:   flag.Uint64(name, defaultValue, usage),
		minimum: minimum,
		maximum: maximum,
	}), err

}
