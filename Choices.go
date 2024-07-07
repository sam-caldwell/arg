package arg

// Choices - Represents a commandline argument which must be part of a list of choices.
//
//	Given a string, verify that the string is one of an enumerated
//	list of choices.
type Choices struct {
	value   *string
	choices []string
}
