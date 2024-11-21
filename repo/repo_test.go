package repo_test

import (
	"testing"

	"github.com/cooldev900/go-rest-concurrency/mock_repo"
	"github.com/cooldev900/go-rest-concurrency/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	// Initialize GoMock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock instance of the repository
	repo := mock_repo.NewMockRepo(ctrl)

	// Define test input and expected output
	item := models.Item{
		ProductID: "1",
		Amount:    1,
	}
	expectedOrder := &models.Order{
		ID:     "123",
		Item:   item,
		Status: models.OrderStatus_Completed,
		Total:  func() *float64 { val := 10.0; return &val }(), // Example Total value
	}

	// Set expectation for the CreateOrder method
	repo.EXPECT().CreateOrder(item).Return(expectedOrder, nil)

	// Call the method on the mock
	order, err := repo.CreateOrder(item)

	// Validate the result
	assert.NoError(t, err, "Expected no error from CreateOrder")
	assert.Equal(t, expectedOrder, order, "Returned order does not match expected order")
}
