package di

import (
	"fmt"
	"reflect"
	"sync"
)

type dependency struct {
	ctorArgs []reflect.Type
	ctorVal  constructorValue
}
type Container struct {
	dependencyMap map[reflect.Type]dependency
	objects       []getObjectFunc
	mu            sync.RWMutex
}

func (c *Container) Register(t any, ctor any, s Scope) error {
	registerType := reflect.TypeOf(t)
	ctorType := reflect.TypeOf(ctor)

	err := validateConstructor(registerType, ctorType)
	if err != nil {
		return fmt.Errorf("failed to validate constructor: %w", err)
	}

	return nil
}
