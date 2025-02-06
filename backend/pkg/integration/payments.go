package integration

import (
	"os"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
)

func ProcessPayment(amount int64, currency, source string) (string, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
	}
	params.SetSource(source)
	ch, err := charge.New(params)
	if err != nil {
		return "", err
	}
	return ch.ID, nil
}