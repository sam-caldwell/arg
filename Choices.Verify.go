package arg

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Verify - Ensure that the given value exists in the set of choices or return error.
//
//	A choice is a string in a set of strings, and the choice must be
//	in this set or an InvalidInput error must be returned.
func (b *Choices) Verify() error {

	for _, valid := range b.choices {

		if *b.value == valid {

			return nil

		}

	}

	return fmt.Errorf(errors.InvalidInput)

}
