package di

import "reflect"

type getObjectFunc func() (any, error)

func newPrototypeGetObjectFunc(args []reflect.Value, c constructorValue) getObjectFunc {
	return func() (any, error) {
		return reflect.Value(c).Call(args)[0].Interface(), nil
	}
}

func newSingletonGetObjectFunc(args []reflect.Value, c constructorValue) getObjectFunc {
	result := &reflect.Value(c).Call(args)[0]

	return func() (any, error) {
		return result.Interface(), nil
	}
}

func newPrototypeGetObjectFuncWithError(args []reflect.Value, c constructorValue) getObjectFunc {
	return func() (any, error) {
		res := reflect.Value(c).Call(args)
		return res[0].Interface(), valueToError(res[1])
	}
}

func newSingletonGetObjectFuncWithError(args []reflect.Value, c constructorValue) getObjectFunc {
	res := reflect.Value(c).Call(args)

	return func() (any, error) {
		return (&res[0]).Interface(), valueToError(res[1])
	}
}

func valueToError(v reflect.Value) error {
	if !v.IsValid() || v.IsNil() {
		return nil
	}

	err, _ := v.Interface().(error)
	return err
}
