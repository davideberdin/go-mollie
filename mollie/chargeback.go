package mollie

type ChargeBack struct {
	Client Client
}

func (c *ChargeBack) GetChargeBack() (string, error) {
	return "", nil
}

func (c *ChargeBack) ListChargeBacks() (string, error) {
	return "", nil
}
