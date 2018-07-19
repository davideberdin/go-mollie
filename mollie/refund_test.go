package mollie

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"net/http"
)

func TestClient_CreateRefund(t *testing.T) {

	// Create Object request
	p := &RefundRequest{
		Amount: map[string]string{
			"currency": "EUR",
			"value":    "100.00",
		},
		Description: "Test Refund",
	}

	// Create Object response
	r := &RefundResponse{
		Resource: "refund",
		ID:       "re_4qqhO89gsT",
		Amount: map[string]string{
			"value":    "100.00",
			"currency": "EUR",
		},
		Description: "Test Refund",
		Status:      "pending",
		PaymentID:   "tr_WDqYK6vllg",
	}

	// Set interceptor
	defer gock.Off()
	gock.New(baseAddress).
		Post("payments/tr_WDqYK6vllg/refunds").
		MatchType("json").
		JSON(p).
		Reply(http.StatusCreated).
		JSON(r)

	c := NewClient("api-key", true)

	r, err := c.CreateRefund(p, "tr_WDqYK6vllg")
	assert.Equal(t, err, nil)

	assert.Equal(t, r.Status, "pending")
	assert.Equal(t, r.Resource, "refund")
	assert.Equal(t, r.Description, "Test Refund")
}

func TestClient_GetRefund(t *testing.T) {

	// Create Object response
	r := &RefundResponse{
		Resource: "refund",
		ID:       "re_4qqhO89gsT",
		Amount: map[string]string{
			"value":    "100.00",
			"currency": "EUR",
		},
		Description: "Test Refund",
		Status:      "pending",
		PaymentID:   "tr_WDqYK6vllg",
	}

	// Set interceptor
	defer gock.Off()
	gock.New(baseAddress).
		Get("payments/tr_WDqYK6vllg/refunds/re_4qqhO89gsT").
		Reply(http.StatusCreated).
		JSON(r)

	c := NewClient("api-key", true)

	r, err := c.GetRefund("tr_WDqYK6vllg", "re_4qqhO89gsT")
	assert.Equal(t, err, nil)

	assert.Equal(t, r.Status, "pending")
	assert.Equal(t, r.Resource, "refund")
	assert.Equal(t, r.Description, "Test Refund")
}

func TestClient_CancelRefund(t *testing.T) {

	// Create Object response
	r := &RefundResponse{
		Resource: "refund",
		ID:       "re_4qqhO89gsT",
		Amount: map[string]string{
			"value":    "100.00",
			"currency": "EUR",
		},
		Description: "Test Refund",
		Status:      "pending",
		PaymentID:   "tr_WDqYK6vllg",
	}

	// Set interceptor
	defer gock.Off()
	gock.New(baseAddress).
		Delete("/payments/tr_WDqYK6vllg/refunds/re_4qqhO89gsT").
		Reply(http.StatusCreated).
		JSON(r)

	c := NewClient("api-key", true)

	r, err := c.CancelRefund("tr_WDqYK6vllg", "re_4qqhO89gsT")
	assert.Equal(t, err, nil)

	assert.Equal(t, r.Status, "pending")
	assert.Equal(t, r.Resource, "refund")
	assert.Equal(t, r.Description, "Test Refund")

}

func TestClient_ListAllRefunds(t *testing.T) {

	payment := &RefundResponse{
		Resource: "refund",
		ID:       "re_4qqhO89gsT",
		Amount: map[string]string{
			"value":    "100.00",
			"currency": "EUR",
		},
		Description: "Test Refund",
		Status:      "pending",
		PaymentID:   "tr_WDqYK6vllg",
	}

	refunds := make([]RefundResponse, 1)

	refunds = append(refunds, *payment)

	// Create Object response
	r := &RefundsResponse{
		Count: 1,
		EmbeddedRefunds: EmbeddedRefunds{
			Refunds: refunds,
		},
		Links: nil,
	}

	// Set interceptor
	defer gock.Off()
	gock.New(baseAddress).
		Get("refunds").
		Reply(http.StatusOK).
		JSON(r)

	c := NewClient("api-key", true)

	r, err := c.ListAllRefunds(nil)

	assert.Equal(t, err, nil)
	assert.True(t, r.Count > 0)
}

func TestClient_ListRefundsOfPayment(t *testing.T) {
	payment := &RefundResponse{
		Resource: "refund",
		ID:       "re_4qqhO89gsT",
		Amount: map[string]string{
			"value":    "100.00",
			"currency": "EUR",
		},
		Description: "Test Refund",
		Status:      "pending",
		PaymentID:   "tr_WDqYK6vllg",
	}

	refunds := make([]RefundResponse, 1)

	refunds = append(refunds, *payment)

	// Create Object response
	r := &RefundsResponse{
		Count: 1,
		EmbeddedRefunds: EmbeddedRefunds{
			Refunds: refunds,
		},
		Links: nil,
	}

	// Set interceptor
	defer gock.Off()
	gock.New(baseAddress).
		Get("payments/tr_WDqYK6vllg/refunds").
		Reply(http.StatusOK).
		JSON(r)

	c := NewClient("api-key", true)

	r, err := c.ListRefundsOfPayment(nil, "tr_WDqYK6vllg")

	assert.Equal(t, err, nil)
	assert.True(t, r.Count > 0)
}
