package entity

import (
	"encoding/json"
	"mohhefni/go-online-shop/apps/transaction/request"
	"mohhefni/go-online-shop/infra/errorpkg"
	"time"
)

type TransactionStatus uint8

const (
	TransactionStatus_Created     TransactionStatus = 1
	TransactionStatus_Progres     TransactionStatus = 10
	TransactionStatus_In_Delivery TransactionStatus = 15
	TransactionStatus_Completed   TransactionStatus = 20

	Trx_Created     = "CREATED"
	Trx_Progres     = "IN PROGRES"
	Trx_In_Delivery = "IN DELIVERY"
	Trx_Completed   = "COMPLITED"

	Trx_Unknown = "UNKNOWN STATUS"
)

var (
	MappintTransactionStatus = map[TransactionStatus]string{
		TransactionStatus_Created:     Trx_Created,
		TransactionStatus_Progres:     Trx_Progres,
		TransactionStatus_In_Delivery: Trx_In_Delivery,
		TransactionStatus_Completed:   Trx_Completed,
	}
)

type TransactionEntity struct {
	Id              string            `db:"id"`
	UserPublicId    string            `db:"user_public_id"`
	ProductId       uint              `db:"product_id"`
	ProductPrice    uint              `db:"product_price"`
	Amount          uint16            `db:"amount"`
	SubTotal        uint              `db:"sub_total"`
	PlatformFee     uint              `db:"platform_fee"`
	GrandTotal      uint              `db:"grand_total"`
	Status          TransactionStatus `db:"status"`
	ProductSnapshot json.RawMessage   `db:"product_snapshot"`
	CreatedAt       time.Time         `db:"created_at"`
	UpdatedAt       time.Time         `db:"updated_at"`
}

func (t *TransactionEntity) ValidateAmount() (err error) {
	if t.Amount == 0 {
		err = errorpkg.ErrAmountInvalid
		return
	}

	return
}

func (t *TransactionEntity) ValidateStock(stock uint16) (err error) {
	if t.Amount > stock {
		err = errorpkg.ErrAmountGreaterThanStock
		return
	}

	return
}

func NewTransactionFromRequest(req request.AddTransactionPayload) *TransactionEntity {
	return &TransactionEntity{
		UserPublicId: req.UserPublicId,
		Amount:       req.Amount,
		Status:       TransactionStatus_Created,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (t *TransactionEntity) SetSubTotal() *TransactionEntity {
	if t.SubTotal == 0 {
		t.SubTotal = t.ProductPrice * uint(t.Amount)
	}

	return t
}

func (t *TransactionEntity) SetPlatformFee(platformFee uint) *TransactionEntity {
	t.PlatformFee = platformFee

	return t
}

// set id, price and call function to  set product snapshot
func (t *TransactionEntity) FromProductToTransaction(product ProductJsonEntity) *TransactionEntity {
	t.ProductId = uint(product.Id)
	t.ProductPrice = uint(product.Price)

	t.SetProductJson(product)

	return t
}

func (t *TransactionEntity) SetGrandTotal() *TransactionEntity {
	if t.GrandTotal == 0 {
		t.SetSubTotal()
		t.GrandTotal = t.SubTotal + t.PlatformFee
	}

	return t
}

// set snapshot
func (t *TransactionEntity) SetProductJson(product ProductJsonEntity) (err error) {
	json, err := json.Marshal(product)
	if err != nil {
		return
	}

	t.ProductSnapshot = json

	return
}

func (t *TransactionEntity) GetProductJson() (product ProductJsonEntity, err error) {
	err = json.Unmarshal(t.ProductSnapshot, &product)
	if err != nil {
		return
	}

	return
}

// get data status in string
func (t *TransactionEntity) GetStatus() string {
	status, ok := MappintTransactionStatus[t.Status]
	if !ok {
		status = Trx_Unknown
	}

	return status
}
