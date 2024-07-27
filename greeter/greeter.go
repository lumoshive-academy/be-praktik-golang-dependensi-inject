package greeter

type Greeter struct {
	Error   bool
	Message string
}

func NewGreeter(name string) *Greeter {
	return &Greeter{Message: "Hello " + name}
}
