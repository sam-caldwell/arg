package arg

import "flag"

// NewFilename - Create a new Filename cli argument
func NewFilename(name string, exists ExistsOrNot, usage string) (o *Filename, err error) {

	if err = validName(&name); err != nil {
		return nil, err
	}

	return &(Filename{
		value:  flag.String(name, "", usage),
		exists: exists,
	}), err

}
