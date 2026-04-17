# Go-DI-Container

Простой DI-контейнер на Go с поддержкой областей видимости и автоматического внедрения зависимостей через аргументы конструктора.

## Возможности

- Регистрация зависимостей через `Register()` и `RegisterWithError()`.
- Автоматическая инъекция аргументов конструктора.
- Поддержка scope:
	- `Singleton` - один объект на контейнер.
	- `Prototype` - новый объект на каждый `GetInstance()`.

## Container Methods

- `NewContainer() *Container`
- `Register(t any, ctor any, s Scope) error`
- `MustRegister(t any, ctor any, s Scope)`
- `RegisterWithError(t any, ctor any, s Scope) error`
- `MustRegisterWithError(t any, ctor any, s Scope)`
- `GetInstance(t any) (any, error)`

## Usage

### 1) Базовый пример: инъекция зависимостей + scope

Файл: [examples/basic/main.go](examples/basic/main.go)

Запуск:

`go run ./examples/basic`

- `Service` создается как `Prototype` (каждый вызов новый объект).
- `Repository` создается как `Singleton` (общий объект между сервисами).
- Зависимость `*Repository` автоматически передается в конструктор `Service`.

### 2) Конструктор с ошибкой

Файл: [examples/with_error/main.go](examples/with_error/main.go)

Запуск:

`go run ./examples/with_error`

- Регистрацию через `RegisterWithError()`.
- Корректную обработку ошибки конструктора.
- Успешное создание объекта после устранения причины ошибки.

## Важно

- Тип при регистрации и тип возращаемый функцией-конструктором должны совпадать.
- Тип при регистрации и тип при `GetInstance()` должны совпадать.
- Для `Register()` конструктор должен возвращать ровно 1 значение: `func(...) T`.
- Для `RegisterWithError()` конструктор должен возвращать 2 значения: `func(...) (T, error)`.

