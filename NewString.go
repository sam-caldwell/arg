package arg

import (
	"flag"
	"regexp"
	"strings"
)

// NewString - Create a new string argument
func NewString(name, defaultValue, usage, validationString string) (o *String, err error) {
	if err = validName(&name); err != nil {
		return nil, err
	}

	return &(String{
		value: flag.String(name, defaultValue, usage),
		validation: func() *regexp.Regexp {
			if strings.TrimSpace(validationString) == "" {
				return nil
			} else {
				return regexp.MustCompile(validationString)
			}
		}(),
	}), err
}
