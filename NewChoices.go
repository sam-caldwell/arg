package arg

import (
	"flag"
)

// NewChoices - Create a new string choice argument.
//
//	Given a string, verify that the string is one of an enumerated
//	list of choices.
func NewChoices(name, defaultValue, usage string, choices ...string) (o *Choices, err error) {

	if err = validName(&name); err != nil {
		return nil, err
	}

	return &(Choices{
		value:   flag.String(name, defaultValue, usage),
		choices: choices,
	}), err

}
