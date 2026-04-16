
# Go-DI-Container

## Design

**Injection Types:**
- **Singleton** — One instance of the object per container.
- **Prototype** — A new instance of the object for each request.

---

## Methods

| Method | Description |
|--------|-------------|
| `Register(ObjectType, ConstructorFunction, InjectionType) error` | Registers a dependency. Returns an error if registration fails. |
| `MustRegister(ObjectType, ConstructorFunction, InjectionType)` | Registers a dependency. Panics if registration fails. |
| `RegisterWithError(ObjectType, ConstructorFunctionWithError, InjectionType) error` | Registers a dependency using a constructor that may return an error. Returns an error if registration fails. |
| `MustRegisterWithError(ObjectType, ConstructorFunctionWithError, InjectionType)` | Registers a dependency using a constructor that may return an error. Panics if registration fails. |
| `GetInstance(ObjectType) (interface{}, error)` | Returns an instance of the object. Returns an error if instance creation fails. |
| `MustGetInstance(ObjectType) interface{}` | Returns an instance of the object. Panics if instance creation fails. |
| `GetInstanceWithError(ObjectType) (interface{}, error)` | Returns an instance of the object using a constructor that may return an error. Returns an error if instance creation fails. |
| `MustGetInstanceWithError(ObjectType) interface{}` | Returns an instance of the object using a constructor that may return an error. Panics if instance creation fails. |

---

### Notes

- `ConstructorFunction` is a function that returns an instance (e.g. `func() *MyType`).
- `ConstructorFunctionWithError` is a function that returns an instance and an error (e.g. `func() (*MyType, error)`).
- Methods with the `Must` prefix panic on any error.

