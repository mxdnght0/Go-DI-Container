package di

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

var (
	ErrObjectNotFound      = errors.New("object not found")
	ErrDependencyNotFound  = errors.New("dependency not found")
	ErrFailedToGetInstance = errors.New("failed to get instance")
)

type dependency struct {
	value         any
	ctorArgs      []reflect.Type
	ctorType      constructorType
	ctorVal       constructorValue
	scope         Scope
	getObjectFunc getObjectFunc
}
type Container struct {
	dependencyMap map[reflect.Type]dependency
	mu            sync.RWMutex
}

func NewContainer() *Container {
	return &Container{
		dependencyMap: make(map[reflect.Type]dependency),
	}
}

func (c *Container) Register(t any, ctor any, s Scope) error {
	return c.register(t, ctor, s, validateConstructor)
}

func (c *Container) MustRegister(t any, ctor any, s Scope) {
	err := c.Register(t, ctor, s)
	if err != nil {
		panic(err)
	}
}

func (c *Container) RegisterWithError(t any, ctor any, s Scope) error {
	return c.register(t, ctor, s, validateConstructorWithError)
}

func (c *Container) MustRegisterWithError(t any, ctor any, s Scope) {
	err := c.register(t, ctor, s, validateConstructorWithError)
	if err != nil {
		panic(err)
	}
}

func (c *Container) register(t any, ctor any, s Scope, validationFunc func(returnType reflect.Type, c constructorType) error) error {
	registerType := reflect.TypeOf(t)
	ctorType := reflect.TypeOf(ctor)
	ctorValue := reflect.ValueOf(ctor)

	err := validationFunc(registerType, ctorType)
	if err != nil {
		return fmt.Errorf("failed to validate constructor: %w", err)
	}

	args := make([]reflect.Type, ctorType.NumIn())
	i := 0
	for arg := range ctorType.Ins() {
		args[i] = arg
		i++
	}

	c.mu.Lock()
	c.dependencyMap[registerType] = dependency{
		value:         t,
		ctorArgs:      args,
		ctorType:      constructorType(ctorType),
		ctorVal:       constructorValue(ctorValue),
		scope:         s,
		getObjectFunc: nil,
	}
	c.mu.Unlock()

	return nil
}

func (c *Container) GetInstance(t any) (any, error) {
	c.mu.RLock()
	d, found := c.dependencyMap[reflect.TypeOf(t)]
	c.mu.RUnlock()
	if !found {
		return nil, fmt.Errorf("%w: %s", ErrObjectNotFound, reflect.TypeOf(t))
	}

	if d.getObjectFunc != nil {
		return d.getObjectFunc()
	}

	args := d.ctorArgs
	valueArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		obj, found := c.dependencyMap[arg]
		if !found {
			return nil, fmt.Errorf("%w: %s", ErrDependencyNotFound, arg)
		}
		valueArg, err := c.GetInstance(obj.value)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrFailedToGetInstance, err)
		}
		valueArgs[i] = reflect.ValueOf(valueArg)
	}

	if d.ctorType.NumOut() == 2 {
		if d.scope == Prototype {
			d.getObjectFunc = newPrototypeGetObjectFunc(valueArgs, d.ctorVal)
		} else {
			d.getObjectFunc = newSingletonGetObjectFunc(valueArgs, d.ctorVal)
		}
	} else {
		if d.scope == Prototype {
			d.getObjectFunc = newPrototypeGetObjectFuncWithError(valueArgs, d.ctorVal)
		} else {
			d.getObjectFunc = newSingletonGetObjectFuncWithError(valueArgs, d.ctorVal)
		}
	}

	return d.getObjectFunc()
}
