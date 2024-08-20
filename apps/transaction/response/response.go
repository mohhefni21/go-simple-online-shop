package response

import (
	"mohhefni/go-online-shop/apps/transaction/entity"
	"time"
)

type TransactionHistoryResponse struct {
	Id           string    `json:"id"`
	UserPublicId string    `json:"user_public_id"`
	ProductId    uint      `json:"product_id"`
	ProductPrice uint      `json:"product_price"`
	Amount       uint16    `json:"amount"`
	SubTotal     uint      `json:"sub_total"`
	PlatformFee  uint      `json:"platform_fee"`
	GrandTotal   uint      `json:"grand_total"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Product entity.ProductJsonEntity `json:"product"`
}

func NewListTransactionHistroyResponse(transactions []entity.TransactionEntity) (transactionList []TransactionHistoryResponse) {
	transactionList = []TransactionHistoryResponse{}

	for _, transaction := range transactions {
		product, err := transaction.GetProductJson()
		if err != nil {
			product = entity.ProductJsonEntity{}
		}

		transactionList = append(transactionList, TransactionHistoryResponse{
			Id:           transaction.Id,
			UserPublicId: transaction.UserPublicId,
			ProductId:    transaction.ProductId,
			ProductPrice: transaction.ProductPrice,
			Amount:       transaction.Amount,
			SubTotal:     transaction.SubTotal,
			PlatformFee:  transaction.PlatformFee,
			GrandTotal:   transaction.GrandTotal,
			Status:       transaction.GetStatus(),
			CreatedAt:    transaction.CreatedAt,
			UpdatedAt:    transaction.UpdatedAt,
			Product:      product,
		})
	}

	return
}
