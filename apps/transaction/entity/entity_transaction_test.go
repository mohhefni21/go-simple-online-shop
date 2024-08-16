package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSubTotal(t *testing.T) {
	t.Run("should return subtotal correctly", func(t *testing.T) {
		// Arrange
		var transaction = TransactionEntity{
			ProductPrice: 10000,
			Amount:       10,
		}
		expected := uint(100000)

		// Action
		transaction.SetSubTotal()

		// Assert
		require.Equal(t, expected, transaction.SubTotal)
	})
}

func TestGrandTotal(t *testing.T) {
	t.Run("should return grandtotal correctly without subtotal", func(t *testing.T) {
		// Arrange
		var transaction = TransactionEntity{
			ProductPrice: 10000,
			Amount:       10,
		}
		expected := uint(100000)

		// Action
		transaction.SetGrandTotal()

		// Assert
		require.Equal(t, expected, transaction.GrandTotal)
	})

	t.Run("should return grandtotal correctly without platform fee", func(t *testing.T) {
		// Arrange
		var transaction = TransactionEntity{
			ProductPrice: 10000,
			Amount:       10,
		}
		expected := uint(100000)

		// Action
		transaction.SetSubTotal()
		transaction.SetGrandTotal()

		// Assert
		require.Equal(t, expected, transaction.GrandTotal)
	})

	t.Run("should return grandtotal correctly with platform fee", func(t *testing.T) {
		// Arrange
		var transaction = TransactionEntity{
			ProductPrice: 10000,
			Amount:       10,
			PlatformFee:  2000,
		}
		expected := uint(102000)

		// Action
		transaction.SetSubTotal()
		transaction.SetGrandTotal()

		// Assert
		require.Equal(t, expected, transaction.GrandTotal)
	})
}

func TestProductSnapshotJson(t *testing.T) {
	t.Run("should not return error and marshal json correcly", func(t *testing.T) {
		// Arrange
		var product = ProductJsonEntity{
			Id:    1,
			SKU:   uuid.NewString(),
			Name:  "shampo",
			Price: 10000,
		}

		transaction := TransactionEntity{}

		// Action
		err := transaction.SetProductJson(product)

		// Arrange
		require.Nil(t, err)
		require.NotNil(t, transaction.ProductSnapshot)
	})

	t.Run("should not return error and unmarshal json correcly", func(t *testing.T) {
		// Arrange
		var productPayload = ProductJsonEntity{
			Id:    1,
			SKU:   uuid.NewString(),
			Name:  "shampo",
			Price: 10000,
		}

		transaction := TransactionEntity{}

		// Action
		err := transaction.SetProductJson(productPayload)
		if err != nil {
			t.Fatal(err)
		}
		product, err := transaction.GetProductJson()

		// Arrange
		require.Nil(t, err)
		require.Equal(t, productPayload, product)
	})
}

func TestStatusTransaction(t *testing.T) {
	// Arrange
	type transactionTest struct {
		name     string
		trx      TransactionEntity
		expected string
	}

	dataTests := []transactionTest{
		{
			name:     "should return status transaction created correctly",
			trx:      TransactionEntity{Status: TransactionStatus_Created},
			expected: Trx_Created,
		},
		{
			name:     "should return status transaction progres correctly",
			trx:      TransactionEntity{Status: TransactionStatus_Progres},
			expected: Trx_Progres,
		},
		{
			name:     "should return status transaction in delivery correctly",
			trx:      TransactionEntity{Status: TransactionStatus_In_Delivery},
			expected: Trx_In_Delivery,
		},
		{
			name:     "should return status transaction completed correctly",
			trx:      TransactionEntity{Status: TransactionStatus_Completed},
			expected: Trx_Completed,
		},
	}

	for _, dataTest := range dataTests {
		t.Run(dataTest.name, func(t *testing.T) {

			// Action and Assert
			status := dataTest.trx.GetStatus()

			require.Equal(t, dataTest.expected, status)
		})
	}
}
