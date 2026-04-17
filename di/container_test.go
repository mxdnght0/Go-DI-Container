package di_test

import (
	"errors"
	"testing"

	"github.com/mxdnght0/Go-DI-Container/di"

	"github.com/stretchr/testify/require"
)

func TestPrototypeScope(t *testing.T) {
	// Arrange
	type Bar struct{ Y int }
	c := di.NewContainer()

	// Act
	err := c.Register(&Bar{}, func() *Bar { return &Bar{Y: 1} }, di.Prototype)
	inst1, err1 := c.GetInstance(&Bar{})
	inst2, err2 := c.GetInstance(&Bar{})

	// Assert
	require.NoError(t, err)
	require.NoError(t, err1)
	require.NoError(t, err2)
	require.NotSame(t, inst1.(*Bar), inst2.(*Bar))
}

func TestSingletonScope(t *testing.T) {
	// Arrange
	type Bar struct{ Y int }
	c := di.NewContainer()

	// Act
	err := c.Register(&Bar{}, func() *Bar { return &Bar{Y: 1} }, di.Singleton)
	inst1, _ := c.GetInstance(&Bar{})
	inst2, _ := c.GetInstance(&Bar{})

	// Assert
	require.NoError(t, err)
	require.Same(t, inst1.(*Bar), inst2.(*Bar))
}

func TestDependencyInjectionInjectsConstructorArgs(t *testing.T) {
	// Arrange
	type Repository struct{ Name string }
	type Service struct{ Repo *Repository }
	c := di.NewContainer()

	errRepo := c.Register(&Repository{}, func() *Repository { return &Repository{Name: "main"} }, di.Singleton)
	errService := c.Register(&Service{}, func(repo *Repository) *Service { return &Service{Repo: repo} }, di.Singleton)

	// Act
	inst, errGet := c.GetInstance(&Service{})
	svc, ok := inst.(*Service)

	// Assert
	require.NoError(t, errRepo)
	require.NoError(t, errService)
	require.NoError(t, errGet)
	require.True(t, ok)
	require.NotNil(t, svc.Repo)
	require.Equal(t, "main", svc.Repo.Name)
}

func TestDependencyInjectionReturnsErrorWhenDependencyMissing(t *testing.T) {
	// Arrange
	type Repository struct{ Name string }
	type Service struct{ Repo *Repository }
	c := di.NewContainer()

	errRegister := c.Register(&Service{}, func(repo *Repository) *Service { return &Service{Repo: repo} }, di.Singleton)

	// Act
	_, errGet := c.GetInstance(&Service{})

	// Assert
	require.NoError(t, errRegister)
	require.Error(t, errGet)
	require.ErrorIs(t, errGet, di.ErrDependencyNotFound)
}

func TestRegisterWithErrorReturnsObjectWhenConstructorSucceeds(t *testing.T) {
	// Arrange
	type Client struct{ ID int }
	c := di.NewContainer()

	errRegister := c.RegisterWithError(&Client{}, func() (*Client, error) {
		return &Client{ID: 10}, nil
	}, di.Prototype)

	// Act
	inst, errGet := c.GetInstance(&Client{})
	client, ok := inst.(*Client)

	// Assert
	require.NoError(t, errRegister)
	require.NoError(t, errGet)
	require.True(t, ok)
	require.Equal(t, 10, client.ID)
}

func TestRegisterWithErrorPropagatesConstructorError(t *testing.T) {
	// Arrange
	type Client struct{ ID int }
	expectedErr := errors.New("constructor failed")
	c := di.NewContainer()

	errRegister := c.RegisterWithError(&Client{}, func() (*Client, error) {
		return nil, expectedErr
	}, di.Prototype)

	// Act
	_, errGet := c.GetInstance(&Client{})

	// Assert
	require.NoError(t, errRegister)
	require.Error(t, errGet)
	require.ErrorIs(t, errGet, expectedErr)
}
