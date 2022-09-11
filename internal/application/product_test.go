package application_test

import (
	"github.com/joaosouzadev/go-hexagonal-arch/internal/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "GeForce RTX 3060"
	product.Active = false
	product.Price = 1000

	err := product.Enable()
	require.Nil(t, err)
	require.True(t, product.Active)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "price must be greater than zero to enable a product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "GeForce RTX 3060"
	product.Active = true
	product.Price = 1000

	err := product.Disable()
	require.Equal(t, "price must be zero to disable a product", err.Error())

	product.Price = 0
	err = product.Disable()
	require.Nil(t, err)
	require.False(t, product.Active)
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "GeForce RTX 3060"
	product.Active = true
	product.Price = -1000

	valid, err := product.IsValid()
	require.False(t, valid)
	require.Equal(t, "product price cannot be less than zero", err.Error())
}

// TODO: other tests
