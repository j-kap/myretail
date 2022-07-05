//go:build integration
// +build integration

package integration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	APP_URL       string
	KNOWN_PRODUCT string
)

func init() {
	APP_URL = os.Getenv("APP_URL")
	KNOWN_PRODUCT = "13860428"
}

func Test_GetProduct(t *testing.T) {
	assert := assert.New(t)

	result := httpGet(assert, KNOWN_PRODUCT, http.StatusOK)

	assert.Contains(result["name"], "Big Lebowski")

	curPrice, ok := result["current_price"].(map[string]interface{})
	assert.True(ok)
	keys := make([]string, 0, len(curPrice))
	for k := range curPrice {
		keys = append(keys, k)
	}
	assert.Contains(keys, "currency_code")
	assert.Contains(keys, "value")

	httpGet(assert, "FAKEID", http.StatusNotFound)
}

func TestPutProduct(t *testing.T) {
	assert := assert.New(t)

	VALUE := "1.99"
	CURRENCY_CODE := "USD"
	putBody := fmt.Sprintf(`{"id":"%s","current_price":{"value":"%s","currency_code":"%s"}}`,
		KNOWN_PRODUCT, VALUE, CURRENCY_CODE)

	result := httpPut(assert, KNOWN_PRODUCT, putBody, http.StatusOK)

	curPrice, ok := result["current_price"].(map[string]interface{})
	assert.True(ok)

	value, ok := curPrice["value"].(string)
	assert.True(ok)
	assert.Equal(value, VALUE)

	code, ok := curPrice["currency_code"].(string)
	assert.True(ok)
	assert.Equal(code, CURRENCY_CODE)

	result = httpGet(assert, KNOWN_PRODUCT, http.StatusOK)
	curPrice, ok = result["current_price"].(map[string]interface{})
	assert.True(ok)

	value, ok = curPrice["value"].(string)
	assert.True(ok)
	assert.Equal(value, VALUE)

	code, ok = curPrice["currency_code"].(string)
	assert.True(ok)
	assert.Equal(code, CURRENCY_CODE)
}

func httpGet(assert *assert.Assertions, productID string, expectedCode int) map[string]interface{} {
	req, err := http.NewRequest(http.MethodGet, reqURL(productID), nil)
	assert.Nil(err)

	client := http.Client{}
	resp, err := client.Do(req)

	assert.Nil(err)
	assert.Equal(expectedCode, resp.StatusCode)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(err)

	var result map[string]interface{}

	err = json.Unmarshal(body, &result)
	assert.Nil(err)

	return result
}

func httpPut(assert *assert.Assertions, productID, putBody string, expectedCode int) map[string]interface{} {
	req, err := http.NewRequest(http.MethodPut, reqURL(productID), strings.NewReader(putBody))
	assert.Nil(err)

	client := http.Client{}
	resp, err := client.Do(req)

	assert.Nil(err)
	assert.Equal(expectedCode, resp.StatusCode)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(err)

	var result map[string]interface{}

	err = json.Unmarshal(body, &result)
	assert.Nil(err)
	return result
}

func reqURL(productID string) string {
	return fmt.Sprintf("%s/products/%s", APP_URL, productID)
}
