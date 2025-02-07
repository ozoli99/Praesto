package payments

import (
	"fmt"
	"time"
)

type Invoice struct {
	InvoiceNumber string    `json:"invoice_number"`
	TransactionID string    `json:"transaction_id"`
	Amount        int64     `json:"amount"`
	Currency      string    `json:"currency"`
	Date          time.Time `json:"date"`
	Description   string    `json:"description"`
}

func GenerateInvoice(transactionID string, amount int64, currency, description string) *Invoice {
	invoiceNumber := fmt.Sprintf("INV-%d", time.Now().UnixNano())
	return &Invoice{
		InvoiceNumber: invoiceNumber,
		TransactionID: transactionID,
		Amount:        amount,
		Currency:      currency,
		Date:          time.Now(),
		Description:   description,
	}
}