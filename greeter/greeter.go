package greeter

type Greeter struct {
	Message string
}

func NewGreeter() *Greeter {
	return &Greeter{Message: "Hello"}
}
