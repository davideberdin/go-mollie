package examples

import (
	"github.com/davideberdin/go-mollie/mollie"
	"fmt"
)

func examplePayments() {
	c := mollie.NewClient("your-api-key", true)

	p := &mollie.PaymentRequest{
		Amount: mollie.Amount{
			Currency: "EUR",
			Value:    "100.00",
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
