package arg

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Verify - Ensure the argument value (integer) is inclusively within the range min,...,max.
func (b *String) Verify() error {
	if b.value == nil {
		return fmt.Errorf("command-line arguments have not been parsed")
	}
	if b.validation == nil || b.validation.MatchString(*b.value) {
		return nil //no validation needed
	} else {
		return fmt.Errorf(errors.InvalidInput+errors.Details, "invalid string")
	}

}
