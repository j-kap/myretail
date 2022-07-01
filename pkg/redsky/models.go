package redsky

type Description struct {
	Title string `json:"title"`
}

type Item struct {
	Description Description `json:"product_description"`
}

type Product struct {
	TCIN string `json:"tcin"`
	Item Item   `json:"item"`
}

type ProductData struct {
	Product Product `json:"product"`
}

type ProductResponse struct {
	Data ProductData `json:"data"`
}

type Err struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Errors []Err `json:"errors"`
}
