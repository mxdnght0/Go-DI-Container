package di

import (
	"errors"
	"reflect"
)

type constructorType reflect.Type
type constructorValue reflect.Value

var (
	ErrTypeNotSet        = errors.New("type not set")
	ErrConstructorNotSet = errors.New("constructor not set")
	ErrNotAFunc          = errors.New("constructor is not a function")
	ErrInvalidNumOutputs = errors.New("invalid number of output values")
	ErrInvalidOutputType = errors.New("output type is invalid")
)

func validateConstructor(returnType reflect.Type, c constructorType) error {
	if c.Kind() != reflect.Func {
		return ErrNotAFunc
	}
	if c.NumOut() != 1 {
		return ErrInvalidNumOutputs
	}
	if c.Out(0) != returnType {
		return ErrInvalidOutputType
	}

	return nil
}

func validateConstructorWithError(t reflect.Type, c constructorType) error {
	if t == nil {
		return ErrTypeNotSet
	}

	if c == nil {
		return ErrConstructorNotSet
	}

	k := c.Kind()
	if k != reflect.Func {
		return ErrNotAFunc
	}
	if c.NumOut() != 2 {
		return ErrInvalidNumOutputs
	}
	if c.Out(0) != t || !c.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		return ErrInvalidOutputType
	}

	return nil
}
