package application_test

import (
	"reflect"
	"testing"

	"github.com/andersonigorf/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	id := uuid.NewV4().String()

	product := application.Product{}
	product.ID = id

	require.NotEmpty(t, product.GetID())
	require.Equal(t, product.GetID(), id)
	require.EqualValues(t, product.GetID(), id)
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "GoMock"

	require.NotEmpty(t, product.GetName())
	require.Equal(t, product.GetName(), "GoMock")
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = application.ENABLED

	require.NotEmpty(t, product.GetStatus())
	require.Equal(t, product.GetStatus(), application.ENABLED)

	product.Status = application.DISABLED
	require.Equal(t, product.GetStatus(), application.DISABLED)
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 10000

	require.NotEmpty(t, product.GetPrice())

	typeOfPrice := reflect.TypeOf(product.Price).Kind()
	require.True(t, typeOfPrice == reflect.Float64)

	require.EqualValues(t, product.GetPrice(), 10000)

}
