package utils_test

import (
	"os"
	"sync"
	"testing"

	"github.com/cooldev900/go-rest-concurrency/models"
	"github.com/cooldev900/go-rest-concurrency/utils"
	"github.com/stretchr/testify/assert"
)

func TestImportProducts(t *testing.T) {
	// Create a temporary CSV file for testing
	tempFile, err := os.CreateTemp("", "products_*.csv")
	assert.NoError(t, err, "Failed to create temporary test file")
	defer os.Remove(tempFile.Name()) // Clean up after test

	// Write test data to the file
	testData := `id,name,stock,category,price
1,ProductA,10,CategoryA,15.50
2,ProductB,5,CategoryB,20.00
3,ProductC,25,CategoryC,30.75
`
	_, err = tempFile.WriteString(testData)
	assert.NoError(t, err, "Failed to write test data to temporary file")

	// Prepare a sync.Map for storing imported products
	products := &sync.Map{}

	// Override the productInputPath with the temporary file path
	oldPath := utils.ProductInputPath
	utils.ProductInputPath = tempFile.Name()
	defer func() { utils.ProductInputPath = oldPath }() // Restore the old path after the test

	// Call the function and verify the output
	err = utils.ImportProducts(products)
	assert.NoError(t, err, "ImportProducts returned an error")

	// Validate the imported products
	expectedProducts := map[string]models.Product{
		"1": {ID: "1", Name: "ProductA(CategoryA)", Stock: 10, Price: 15.50},
		"2": {ID: "2", Name: "ProductB(CategoryB)", Stock: 5, Price: 20.00},
		"3": {ID: "3", Name: "ProductC(CategoryC)", Stock: 25, Price: 30.75},
	}

	products.Range(func(key, value interface{}) bool {
		id := key.(string)
		product := value.(models.Product)

		expected, exists := expectedProducts[id]
		assert.True(t, exists, "Unexpected product ID found: %s", id)
		assert.Equal(t, expected, product, "Product data mismatch for ID: %s", id)
		return true
	})
}
