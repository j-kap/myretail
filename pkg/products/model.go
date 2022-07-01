package products

type Product struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price Price  `json:"current_price,omitempty"`
}

type Price struct {
	Value    string `json:"value,omitempty"`
	Currency string `json:"currency_code,omitempty"`
}
