package arg

import (
	"fmt"
	"github.com/sam-caldwell/errors"
	"regexp"
)

// validName - ensure the argument name is valid
func validName(name *string) error {

	const expression = `[a-zA-Z][a-zA-Z0-9\-_.]*`

	if pattern := regexp.MustCompile(expression); pattern.MatchString(*name) {
		return nil
	}

	return fmt.Errorf(errors.InvalidInput)

}
