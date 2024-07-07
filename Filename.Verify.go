package arg

import (
	"fmt"
	"github.com/sam-caldwell/errors"
	"github.com/sam-caldwell/file"
	"strings"
)

// Verify - Ensure the given filename is a valid filename and (optionally) that it exists.
func (b *Filename) Verify() error {

	if strings.TrimSpace(*b.value) == "" {
		return fmt.Errorf(errors.InvalidInput+errors.Details, "empty filename")
	}

	if b.exists == Exists && !file.Exists(*b.value) {
		return fmt.Errorf(errors.NotFound)
	}

	return nil

}
