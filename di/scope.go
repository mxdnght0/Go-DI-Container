package di

type Scope int

const (
	Prototype Scope = iota
	Singleton
)
