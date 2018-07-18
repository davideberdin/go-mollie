package mollie

import (
	"github.com/dghubble/sling"
)

// TODO: Extend this file with OAuth
// For now let's keep it simple

func NewClient(apiKey string, testMode bool) (Client) {
	c := Client{
		sling:    sling.New(),
		baseURL:  baseAddress,
		apiKey:   apiKey,
		TestMode: testMode,
	}

	// set authorization for every request
	c.sling.Base(baseAddress).Set("Authorization", "Bearer "+c.apiKey)

	return c
}
