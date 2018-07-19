package main

import (
	"fmt"
	"github.com/davideberdin/go-mollie/mollie"
	"github.com/spf13/viper"
	"log"
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

func main() {
	c := mollie.NewClient("your-api-key", true)

	p := &mollie.PaymentRequest{
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
		panic(err)
	}

	fmt.Println(r)
}


