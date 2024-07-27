package service

import "golang-dependensi-inject/greeter"

type Service struct {
	Greeter *greeter.Greeter
}

func NewService(greeter *greeter.Greeter) *Service {
	return &Service{Greeter: greeter}
}
