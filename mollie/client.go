package mollie

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

type Client struct {
	ApiKey   string
	BaseURL  string
	TestMode bool
}

func NewClient(apiKey string, testMode bool) (*Client) {
	c := &Client{
		ApiKey:   apiKey,
		BaseURL:  "https://api.mollie.com/v2/",
		TestMode: testMode,
	}
	return c
}

func (c *Client) GetRequest(endpoint string) (string, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c *Client) PostRequest(endpoint string) (string, error) {

	return "", nil
}

func (c *Client) DeleteRequest(endpoint string) (string, error) {

	return "", nil
}
