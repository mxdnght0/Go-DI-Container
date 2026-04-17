package main

import (
	"errors"
	"fmt"

	"github.com/mxdnght0/Go-DI-Container/di"
)

type DBClient struct {
	DSN string
}

func main() {
	c := di.NewContainer()
	shouldFail := true

	err := c.RegisterWithError(&DBClient{}, func() (*DBClient, error) {
		if shouldFail {
			return nil, errors.New("database is unavailable")
		}
		return &DBClient{DSN: "postgres://localhost:5432/app"}, nil
	}, di.Prototype)
	if err != nil {
		panic(err)
	}

	_, err = c.GetInstance(&DBClient{})
	fmt.Println("first attempt failed:", err != nil)
	if err != nil {
		fmt.Println("error:", err)
	}

	shouldFail = false
	inst, err := c.GetInstance(&DBClient{})
	if err != nil {
		panic(err)
	}

	client := inst.(*DBClient)
	fmt.Println("second attempt succeeded:", client.DSN)
}
