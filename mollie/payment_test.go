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

func TestClient_CreatePayment(t *testing.T) {
	apiKey := configuration()
	c := NewClient(apiKey, true)

	p := &PaymentRequest{
		Amount: map[string]string{
			"currency": "EUR",
			"value":    "100.00",
		},
		Description: "Testing Payment",
		RedirectURL: "http://2e9fafad.ngrok.io",
		WebhookURL:  "http://2e9fafad.ngrok.io",
		Method:      "creditcard",
		BillingAddress: Address{
			StreetAndNumber: "Gustav Mahlerlaand 869",
			PostalCode:      "1082MK",
			City:            "Amsterdam",
			Region:          "Noord Holland",
			Country:         "NL",
		},
	}

	r, err := c.CreatePayment(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	assert.Equal(t, r.Status, 201)
}
