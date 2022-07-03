//go:build integration
// +build integration

package integration

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func Test_GetProduct(t *testing.T) {
	APP_URL := os.Getenv("APP_URL")
	assert := assert.New(t)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/products/13860428", APP_URL), nil)
	assert.Nil(err)

	client := http.Client{}
	resp, err := client.Do(req)

	assert.Nil(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(err)

	var result map[string]interface{}

	err = json.Unmarshal(body, &result)
	assert.Nil(err)

	assert.Contains(result["name"], "Big Lebowski")
}
