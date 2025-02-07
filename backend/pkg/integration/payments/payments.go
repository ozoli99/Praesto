package payments

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
	"github.com/stripe/stripe-go/v72/dispute"
	"github.com/stripe/stripe-go/v72/refund"
)

func ProcessPayment(amount int64, currency, source string) (string, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
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

func ProcessRefund(chargeID string, amount int64) (string, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
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

func ListDisputes(limit int) ([]*stripe.Dispute, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	params := &stripe.DisputeListParams{}
	params.Filters.AddFilter("limit", "", fmt.Sprintf("%d", limit))

	i := dispute.List(params)
	var disputes []*stripe.Dispute
	for i.Next() {
		disputes = append(disputes, i.Dispute())
	}
	if err := i.Err(); err != nil {
		return nil, err
	}
	return disputes, nil
}
