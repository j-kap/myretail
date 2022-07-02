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

type Client struct {
	BaseURL string
}

func New(baseURL string) Client {
	return Client{baseURL}
}

func (c Client) GetProduct(id string) (ProductResponse, error) {
	var prod ProductResponse

	client := &http.Client{}
	req, err := http.NewRequest("GET", c.BaseURL, nil)
	if err != nil {
		return prod, err
	}

	req.Header.Add("Accept", "application/json")
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
		err = fmt.Errorf(parseErrors(resp_body))
		return prod, err
	}

	if err = json.Unmarshal(resp_body, &prod); err != nil {
		return prod, err
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
