package payments

import (
	"errors"
)

type PaymentAdapter interface {
	ProcessPayment(amount int64, currency, source string) (string, error)
	ProcessRefund(chargeID string, amount int64) (string, error)
}

func NewPaymentAdapter(adapterFlag, secretKey string) (PaymentAdapter, error) {
	switch adapterFlag {
		case "stripe":
			return NewStripeAdapter(secretKey), nil
		case "paypal":
			return NewPayPalAdapter(secretKey)
		default:
			return nil, errors.New("Unknown payment adapter")
	}
}