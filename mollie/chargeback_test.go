package mollie

import (
	"testing"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetChargeBack(t *testing.T) {

	// Create Object response
	r := &ChargebackResponse{
		Resource: "refund",
		ID:       "chb_n9z0tp",
		Amount: Amount{
			Value:    "100.00",
			Currency: "EUR",
		},
		PaymentID:   "tr_WDqYK6vllg",
		SettlementAmount: Amount{
			Value:    "100.00",
			Currency: "EUR",
		},
	}

	// Set interceptor
	defer gock.Off()
	gock.New(baseAddress).
		Get("payments/tr_WDqYK6vllg/chargebacks/chb_n9z0tp").
		Reply(http.StatusCreated).
		JSON(r)

	c := NewClient("api-key", true)

	r, err := c.GetChargeBack("tr_WDqYK6vllg", "chb_n9z0tp")
	assert.Equal(t, err, nil)

	assert.Equal(t, r.Resource, "refund")
	assert.Equal(t, r.ID, "chb_n9z0tp")
	assert.Equal(t, r.SettlementAmount.Value, "100.00")
}

func TestClient_ListAllChargeBacks(t *testing.T) {

	chargeback := &ChargebackResponse{
		Resource: "refund",
		ID:       "chb_n9z0tp",
		Amount: Amount{
			Value:    "100.00",
			Currency: "EUR",
		},
		PaymentID:   "tr_WDqYK6vllg",
	}

	chargebacks := make([]ChargebackResponse, 1)

	chargebacks = append(chargebacks, *chargeback)

	// Create Object response
	r := &ChargebacskResponse{
		Count: 1,
		EmbeddedChargeback: EmbeddedChargeback{
			Chargeback: chargebacks,
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

	r, err := c.ListAllChargeBacks()

	assert.Equal(t, err, nil)
	assert.True(t, r.Count > 0)
}

func TestClient_ListChargebacksOfPayment(t *testing.T) {
	chargeback := &ChargebackResponse{
		Resource: "refund",
		ID:       "chb_n9z0tp",
		Amount: Amount{
			Value:    "100.00",
			Currency: "EUR",
		},
		PaymentID:   "tr_WDqYK6vllg",
	}

	chargebacks := make([]ChargebackResponse, 1)

	chargebacks = append(chargebacks, *chargeback)

	// Create Object response
	r := &ChargebacskResponse{
		Count: 1,
		EmbeddedChargeback: EmbeddedChargeback{
			Chargeback: chargebacks,
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

	r, err := c.ListChargeBacksOfPayment("tr_WDqYK6vllg")

	assert.Equal(t, err, nil)
	assert.True(t, r.Count > 0)
}
