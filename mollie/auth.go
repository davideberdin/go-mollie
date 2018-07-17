package mollie

import "net/http"

// TODO: Extend this file with OAuth
// For now let's keep it simple

func NewClient(apiKey string, testMode bool) (Client) {
	c := Client{
		http: &http.Client{},
		apiKey:   apiKey,
		baseURL:  baseAddress,
		TestMode: testMode,
	}
	return c
}
