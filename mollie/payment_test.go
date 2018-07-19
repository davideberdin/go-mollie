package mollie

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"net/http"
)

func TestClient_CreatePayment(t *testing.T) {

	// Create Object request
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

	// Create Object response
	r := &PaymentResponse{
		Resource: "payment",
		ID:       "tr_WthPtuq48H",
		Mode:     "test",
		Amount: map[string]string{
			"value":    "100.00",
			"currency": "EUR",
		},
		Description: "Test Payment",
		Method:      "banktransfer",
		Status:      "open",
	}

	// Set interceptor
	defer gock.Off()
	gock.New(baseAddress).
		Post("/payments").
		MatchType("json").
		JSON(p).
		Reply(http.StatusCreated).
		JSON(r)

	c := NewClient("api-key", true)

	r, err := c.CreatePayment(p)
	assert.Equal(t, err, nil)

	assert.Equal(t, r.Status, "open")
	assert.Equal(t, r.Resource, "payment")
	assert.Equal(t, r.Mode, "test")
	assert.Equal(t, r.Description, "Testing Payment")
	assert.Equal(t, r.Method, "banktransfer")
}

func TestClient_GetPayment(t *testing.T) {

	// Create Object response
	r := &PaymentResponse{
		Resource: "payment",
		ID:       "tr_WthPtuq48H",
		Mode:     "test",
		Amount: map[string]string{
			"value":    "100.00",
			"currency": "EUR",
		},
		Description: "Test Payment",
		Method:      "banktransfer",
		Status:      "open",
	}

	// Set interceptor
	defer gock.Off()
	gock.New(baseAddress).
		Get("/payments/tr_WthPtuq48H").
		Reply(http.StatusOK).
		JSON(r)

	c := NewClient("api-key", true)

	r, err := c.GetPayment("tr_WthPtuq48H", nil)

	assert.Equal(t, err, nil)
	assert.Equal(t, r.Status, "open")
	assert.Equal(t, r.ID, "tr_WthPtuq48H")
}

func TestClient_CancelPayment(t *testing.T) {

	// Create Object response
	r := &PaymentResponse{
		Resource: "payment",
		ID:       "tr_WthPtuq48H",
		Mode:     "test",
		Amount: map[string]string{
			"value":    "100.00",
			"currency": "EUR",
		},
		Description: "Test Payment",
		Method:      "banktransfer",
		Status:      "canceled",
	}

	// Set interceptor
	defer gock.Off()
	gock.New(baseAddress).
		Get("/payments/tr_WthPtuq48H").
		Reply(http.StatusOK).
		JSON(r)

	c := NewClient("api-key", true)

	r, err := c.GetPayment("tr_WthPtuq48H", nil)

	assert.Equal(t, err, nil)
	assert.Equal(t, r.Status, "canceled")
	assert.Equal(t, r.ID, "tr_WthPtuq48H")
}

func TestClient_ListPayments(t *testing.T) {

	payment := &PaymentResponse{
		Resource: "payment",
		ID:       "tr_WthPtuq48H",
		Mode:     "test",
		Amount: map[string]string{
			"value":    "100.00",
			"currency": "EUR",
		},
		Description: "Test Payment",
		Method:      "banktransfer",
		Status:      "canceled",
	}

	payments := make([]PaymentResponse, 1)

	payments = append(payments, *payment)

	// Create Object response
	r := &PaymentsResponse{
		Count: 1,
		Embedded: struct {
			Payments []PaymentResponse
		}{
			Payments: payments,
		},
		Links: nil,
	}

	// Set interceptor
	defer gock.Off()
	gock.New(baseAddress).
		Get("/payments").
		Reply(http.StatusOK).
		JSON(r)

	c := NewClient("api-key", true)

	r, err := c.ListPayments(nil)

	assert.Equal(t, err, nil)
	assert.True(t, r.Count > 0)
}
