package main

import (
	"Go-DI-Container/di"
	"fmt"
)

type Repository struct {
	greet string
}

func NewRepository() *Repository {
	return &Repository{
		greet: "Hello World",
	}
}

type Service struct {
	repository *Repository
}

func (s *Service) Greet() {
	fmt.Println("Hello World")
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func main() {
	c := di.NewContainer()
	err := c.Register((*Service)(nil), NewService, di.Prototype)
	if err != nil {
		panic(err)
	}
	err = c.Register((*Repository)(nil), NewRepository, di.Prototype)
	if err != nil {
		panic(err)
	}
	svc, err := c.GetInstance((*Service)(nil))
	if err != nil {
		panic(err)
	}
	s := svc.(*Service)
	s.Greet()
}
