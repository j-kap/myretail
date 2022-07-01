package redsky

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	Err404NotFound = errors.New("Product not found")
)

const BASE_URL = "https://my-products-api.com/api/v1/products?key=super-secret-key"

type client struct{}

func New() client {
	return client{}
}

func (c client) GetProduct(id string) (ProductResponse, error) {
	var prod ProductResponse

	client := &http.Client{}
	req, err := http.NewRequest("GET", BASE_URL, nil)
	req.Header.Add("Accept", "application/json")

	if err != nil {
		return prod, err
	}

	q := req.URL.Query()
	q.Add("tcin", id)

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return prod, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return prod, Err404NotFound
	}

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return prod, err
	}

	if resp.StatusCode != 200 {
		return prod, fmt.Errorf("Unknown error(s) getting product: %s", parseErrors(resp_body))
	}

	if err = json.Unmarshal(resp_body, &prod); err != nil {
		return prod, fmt.Errorf("Error unmarshaling product info: %s", err)
	}

	return prod, nil
}

func parseErrors(body []byte) string {
	var resp ErrorResponse

	err := json.Unmarshal(body, &resp)
	if err != nil {
		return fmt.Sprintf("Unmarshal error %s while reading response: %s", err, body)
	}

	errors := ""
	for _, err := range resp.Errors {
		errors = fmt.Sprintf("%s, %s", errors, err)
	}

	return errors
}
