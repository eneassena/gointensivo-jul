package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Order_If_It_Gets_An_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order{}
	assert.Equal(t, order.Validate().Error(), "id is required")
}

func Test_Order_If_It_Gets_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{ID: "50"}
	assert.Equal(t, order.Validate().Error(), "price is required")
}

func Test_Order_If_It_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{ID: "50", Price: 10}
	assert.Equal(t, order.Validate().Error(), "tax is required")
}

func Test_Order_Create(t *testing.T) {
	order := Order{ID: "50", Price: 10, Tax: 5}
	assert.EqualValues(t, "50", order.ID)
	assert.EqualValues(t, 10, order.Price)
	assert.EqualValues(t, 5, order.Tax)
	order.CalculateFinalPrice()
	assert.EqualValues(t, 15, order.FinalPrice)

}
