package di

import "reflect"

type getObjectFunc func() (any, error)

func newPrototypeGetObjectFunc(args []reflect.Value, c constructorValue) getObjectFunc {
	return func() (any, error) {
		return reflect.Value(c).Call(args)[0], nil
	}
}

func newSingletonGetObjectFunc(args []reflect.Value, c constructorValue) getObjectFunc {
	result := reflect.Value(c).Call(args)[0]

	return func() (any, error) {
		return result, nil
	}
}

func newPrototypeGetObjectFuncWithError(args []reflect.Value, c constructorValue) getObjectFunc {
	return func() (any, error) {
		res := reflect.ValueOf(c).Call(args)
		return res[0], res[1].Interface().(error)
	}
}

func newSingletonGetObjectFuncWithError(args []reflect.Value, c constructorValue) getObjectFunc {
	res := reflect.Value(c).Call(args)

	return func() (any, error) {
		return res[0], res[1].Interface().(error)
	}
}
