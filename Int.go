package arg

// Int - Represents a command-line integer argument (e.g. --foo 1) with bounds checking.
//
//	The bounds checking must assume the input value (v) is within the range min...max,
//	inclusively.
type Int struct {
	value   *int64
	minimum int64
	maximum int64
}
