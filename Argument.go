package arg

type Argument interface {
	Value() interface{}
	Verify() error
}
