package arg

// Uint - Represents a command-line unsigned 64-bit integer argument (e.g. --foo 1) with bounds checking.
//
//	The bounds checking must assume the input value (v)
//	is within the range min...max, inclusively.
type Uint struct {
	value   *uint64
	minimum uint64
	maximum uint64
}
