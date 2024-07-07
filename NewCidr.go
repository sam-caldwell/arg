package arg

import (
	"flag"
)

// NewCidr - Create a new CIDR string argument
func NewCidr(name, defaultValue, usage string) (o *Cidr, err error) {
	if err = validName(&name); err != nil {
		return nil, err
	}

	return &(Cidr{
		value: flag.String(name, defaultValue, usage),
	}), err
}
