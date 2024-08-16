package entity

import (
	"encoding/json"
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
	Email           string            `db:"email"`
	ProductId       uint              `db:"product_id"`
	ProductPrice    uint              `db:"product_price"`
	Amount          uint8             `db:"amount"`
	SubTotal        uint              `db:"sub_total"`
	PlatformFee     uint              `db:"platform_fee"`
	GrandTotal      uint              `db:"grand_total"`
	Status          TransactionStatus `db:"status"`
	ProductSnapshot json.RawMessage   `db:"product_snapshot"`
	CreatedAt       time.Time         `db:"created_at"`
	UpdatedAt       time.Time         `db:"updated_at"`
}

func NewTransactionFromRequest(email string) *TransactionEntity {
	return &TransactionEntity{
		Email:     email,
		Status:    TransactionStatus_Created,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t *TransactionEntity) SetSubTotal() {
	if t.SubTotal == 0 {
		t.SubTotal = t.ProductPrice * uint(t.Amount)
	}
}

func (t *TransactionEntity) FromProductToTransaction(product ProductJsonEntity) {
	t.ProductId = uint(product.Id)
	t.ProductPrice = uint(product.Price)
}

func (t *TransactionEntity) SetGrandTotal() {
	if t.GrandTotal == 0 {
		t.SetSubTotal()
		t.GrandTotal = t.SubTotal + t.PlatformFee
	}
}

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

func (t *TransactionEntity) GetStatus() string {
	status, ok := MappintTransactionStatus[t.Status]
	if !ok {
		status = Trx_Unknown
	}

	return status
}
