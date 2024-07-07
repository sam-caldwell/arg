package arg

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Verify - Ensure the argument value (integer) is inclusively within the range min,...,max.
func (b *Int) Verify() error {

	if *b.value < b.minimum || *b.value > b.maximum {
		return fmt.Errorf(errors.BoundsCheckError)
	}

	return nil

}
