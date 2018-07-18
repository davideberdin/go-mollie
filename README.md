# Mollie Go SDK
Mollie wrapper API written in Go.

```bash
go get github.com/davideberdin/go-mollie/mollie
```

### Usage Example

```
package main

import (
	"fmt"
	"github.com/davideberdin/go-mollie/mollie"
)

func main() {
    // true for testing mode
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
```
