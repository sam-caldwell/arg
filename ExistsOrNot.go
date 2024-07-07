package arg

// ExistsOrNot - type indicates whether an entity exists or not.
type ExistsOrNot bool

const (
	// Exists - The subject entity exists
	Exists ExistsOrNot = true

	// NotExists - The subject entity does not exist
	NotExists ExistsOrNot = false
)
