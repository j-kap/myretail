package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
)

const PricesCollection = "product-prices"

type Client struct {
	client *firestore.Client
}

func New(client *firestore.Client) Client {
	return Client{client}
}

func (c Client) GetProductPrice(ctx context.Context, id string) (ProductPrice, error) {
	var price ProductPrice

	snap, err := c.client.Collection(PricesCollection).Doc(id).Get(ctx)
	if err != nil {
		return price, err
	}

	if err = snap.DataTo(&price); err != nil {
		return price, err
	}

	return price, nil
}

func (c Client) SetProductPrice(ctx context.Context, id, price, currency string) error {
	prodPrice := ProductPrice{
		Value:    price,
		Currency: currency,
	}

	_, err := c.client.Collection(PricesCollection).Doc(id).Set(ctx, prodPrice)
	if err != nil {
		return err
	}

	return nil
}
