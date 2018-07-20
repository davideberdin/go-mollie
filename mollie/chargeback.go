package mollie

import (
	"time"
	"fmt"
)

const chargebackEndpoint = "chargebacks"

// ChargebackResponse defines the objecy for every response from the Mollie APIs regarding chargebacks
// https://docs.mollie.com/reference/v2/chargebacks-api/get-chargeback
type ChargebackResponse struct {
	Resource         string                 `json:"resource"`
	ID               string                 `json:"id"`
	Amount           Amount                 `json:"amount"`
	SettlementAmount Amount                 `json:"settlementAmount"`
	CreatedAt        time.Time              `json:"createdAt"`
	ReversedAt       interface{}            `json:"reversedAt"`
	PaymentID        string                 `json:"paymentId"`
	Links            map[string]interface{} `json:"_links"`
}

func (c *Client) GetChargeBack(paymentId, chargebackId string) (*ChargebackResponse, error) {

	chargebackURL := fmt.Sprintf("%s/%s/%s/%s", paymentsEndpoint, paymentId, chargebackEndpoint, chargebackId)

	var r ChargebackResponse
	err := c.get(chargebackURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type EmbeddedChargeback struct {
	Chargeback []ChargebackResponse `json:"chargebacks"`
}

type ChargebacskResponse struct {
	Count              int                    `json:"count"`
	EmbeddedChargeback EmbeddedChargeback     `json:"_embedded"`
	Links              map[string]interface{} `json:"_links"`
}

func (c *Client) ListAllChargeBacks() (*ChargebacskResponse, error) {
	chargebackURL := fmt.Sprintf("%s", refundsEndpoint)

	var r ChargebacskResponse
	err := c.get(chargebackURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) ListChargeBacksOfPayment(paymentId string) (*ChargebacskResponse, error) {
	chargebackURL := fmt.Sprintf("%s/%s/%s", paymentsEndpoint, paymentId, refundsEndpoint)

	var r ChargebacskResponse
	err := c.get(chargebackURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
