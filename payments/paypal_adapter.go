package payments

import (
	"context"
	"errors"
	"strings"

	paypal "github.com/plutov/paypal/v4"
)

type PayPalAdapter struct {
	Client *paypal.Client
}

func NewPayPalAdapter(secretKey string) (PaymentAdapter, error) {
	parts := strings.SplitN(secretKey, ":", 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid PayPal secretKey format; expected 'clientID:clientSecret'")
	}
	clientID := parts[0]
	clientSecret := parts[1]
	client, err := paypal.NewClient(clientID, clientSecret, paypal.APIBaseSandBox)
	if err != nil {
		return nil, err
	}
	context := context.Background()
	_, err = client.GetAccessToken(context)
	if err != nil {
		return nil, err
	}
	return &PayPalAdapter{Client: client}, nil
}

func (adapter *PayPalAdapter) ProcessPayment(amount int64, currency, source string) (string, error) {
	// A real implementation would use adapter.client.CreateOrder and then capture the order.
	return "PAYPAL_ORDER_ID", nil
}

func (adapter *PayPalAdapter) ProcessRefund(chargeID string, amount int64) (string, error) {
	// This is a simplified example.
	return "PAYPAL_REFUND_ID", nil
}