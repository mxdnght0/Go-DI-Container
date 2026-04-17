package main

import (
	"fmt"

	"github.com/mxdnght0/Go-DI-Container/di"
)

type Repository struct {
	Name string
}

type Service struct {
	Repo *Repository
}

func main() {
	c := di.NewContainer()

	if err := c.Register(&Repository{}, func() *Repository {
		return &Repository{Name: "main-repo"}
	}, di.Singleton); err != nil {
		panic(err)
	}

	if err := c.Register(&Service{}, func(repo *Repository) *Service {
		return &Service{Repo: repo}
	}, di.Prototype); err != nil {
		panic(err)
	}

	inst1, err := c.GetInstance(&Service{})
	if err != nil {
		panic(err)
	}
	inst2, err := c.GetInstance(&Service{})
	if err != nil {
		panic(err)
	}

	s1 := inst1.(*Service)
	s2 := inst2.(*Service)

	fmt.Println("service objects are different:", s1 != s2)
	fmt.Println("repository is shared (singleton):", s1.Repo == s2.Repo)
	fmt.Println("repo name:", s1.Repo.Name)
}
