package payments

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
	"github.com/stripe/stripe-go/v72/refund"
)

type StripeAdapter struct {
	SecretKey string
}

func NewStripeAdapter(secretKey string) PaymentAdapter {
	return &StripeAdapter{SecretKey: secretKey}
}

func (adapter *StripeAdapter) ProcessPayment(amount int64, currency, source string) (string, error) {
	stripe.Key = adapter.SecretKey
	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
	}
	params.SetSource(source)
	charge, err := charge.New(params)
	if err != nil {
		return "", err
	}
	return charge.ID, nil
}

func (adapter *StripeAdapter) ProcessRefund(chargeID string, amount int64) (string, error) {
	stripe.Key = adapter.SecretKey
	params := &stripe.RefundParams{
		Charge: stripe.String(chargeID),
	}
	if amount > 0 {
		params.Amount = stripe.Int64(amount)
	}
	refund, err := refund.New(params)
	if err != nil {
		return "", err
	}
	return refund.ID, nil
}
