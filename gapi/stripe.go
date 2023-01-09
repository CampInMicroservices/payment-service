package gapi

import (
	"errors"

	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/paymentintent"
)

func (server *ServerGRPC) InitialisePaymentIntent(amountCents int64, itemID string) (string, error) {

	stripe.Key = server.config.StripeKey

	// Now, we initialise the payment intent with the relevant data
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amountCents),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	}
	// this is very important because we need the item id for webhooks, etc. as all metadata will be sent back by stripe
	params.AddMetadata("item_id", itemID)

	// this makes a request to stripes api to create the payment intent
	intent, err := paymentintent.New(params)

	// if for some reason this fails, it's from stripe. nothing we can do
	if err != nil {
		return "", errors.New("could not initialise payment. An unexpected error occured")
	}

	// return the intent's client secret. this is used by the frontend
	return intent.ClientSecret, nil
}
