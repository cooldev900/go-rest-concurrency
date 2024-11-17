package repo_test

import (
	"testing"

	"github.com/cooldev900/go-rest-concurrency/models"
	"github.com/cooldev900/go-rest-concurrency/repo"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	item := models.Item{ProductID: "1", Amount: 1}

	r, err := repo.New()
	assert.NoError(t, err, "Failed to create repo instance")

	order, err := r.CreateOrder(item)

	assert.NoError(t, err, "CreateOrder returned an error")
	assert.NotNil(t, order, "CreateOrder returned nil order")
	assert.Equal(t, item.ProductID, order.Item.ProductID, "ProductID does not match")
	assert.Equal(t, item.Amount, order.Item.Amount, "Amount does not match")
	assert.Equal(t, models.OrderStatus_Completed, order.Status, "Order status is not 'Completed'")
	assert.NotNil(t, order.Total, "Order total should not be nil")
	assert.Greater(t, *order.Total, 0.0, "Order total should be greater than 0")
}
