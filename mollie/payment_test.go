package mollie

import (
	"testing"
	"fmt"
	"os"
	"github.com/spf13/viper"
	"log"
	"github.com/stretchr/testify/assert"
)

func configuration() string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var c struct {
		ApiKey string
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return c.ApiKey
}

// TODO: Need to create mock functions without using the actual Mollie APIs
func TestClient_CreatePayment(t *testing.T) {
	apiKey := configuration()
	c := NewClient(apiKey, true)

	p := &PaymentRequest{
		Amount: map[string]string{
			"currency": "EUR",
			"value":    "100.00",
		},
		Description:  "Testing Payment",
		RedirectURL:  "http://2e9fafad.ngrok.io",
		WebhookURL:   "http://2e9fafad.ngrok.io",
		Method:       "banktransfer",
		BillingEmail: "sloth@greeny.com",
		DueDate:      "2018-09-12",
		Locale:       "nl_NL",
	}

	r, err := c.CreatePayment(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	assert.Equal(t, r.Status, "open")
	assert.Equal(t, r.Resource, "payment")
	assert.Equal(t, r.Mode, "test")
	assert.Equal(t, r.Description, "Testing Payment")
	assert.Equal(t, r.Method, "banktransfer")
}

func TestClient_GetPayment(t *testing.T) {
	apiKey := configuration()
	c := NewClient(apiKey, true)

	r, err := c.GetPayment("tr_GfdpAP8xnf", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	assert.Equal(t, r.Status, "open")
}

func TestClient_CancelPayment(t *testing.T) {
	apiKey := configuration()
	c := NewClient(apiKey, true)

	r, err := c.GetPayment("tr_GfdpAP8xnf", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	assert.Equal(t, r.Status, "canceled")
}

func TestClient_ListPayments(t *testing.T) {
	apiKey := configuration()
	c := NewClient(apiKey, true)

	r, err := c.ListPayments(nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	assert.True(t, r.Count > 0)
}