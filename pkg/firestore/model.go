package firestore

type ProductPrice struct {
	ID       string `json:"id"`
	Value    string `json:"value"`
	Currency string `json:"currency_code"`
}
