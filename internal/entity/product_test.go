package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 10.00)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, p.Name, "Product 1")
	assert.Equal(t, p.Price, 10.00)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10.00)
	assert.Nil(t, p)
	assert.Equal(t, err, ErrNameIsRequired)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Product 1", 0)
	assert.Nil(t, p)
	assert.Equal(t, err, ErrPriceIsRequired)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Product 1", -10.00)
	assert.Nil(t, p)
	assert.Equal(t, err, ErrInvalidPrice)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Product 1", 10.00)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
