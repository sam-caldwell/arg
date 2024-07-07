package arg

import (
	"flag"
)

// NewInt - Create a new 64-bit integer flag argument
func NewInt(name string, defaultValue, minimum, maximum int64, usage string) (o *Int, err error) {

	if err = validName(&name); err != nil {
		return nil, err
	}

	return &(Int{
		value:   flag.Int64(name, defaultValue, usage),
		minimum: minimum,
		maximum: maximum,
	}), err

}
