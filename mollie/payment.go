package mollie

type Payment struct {
	Client Client
}

func (p *Payment) CreatePayment() (string, error) {
	return "", nil
}

func (p *Payment) GetPayment() (string, error) {
	return "", nil
}

func (p *Payment) CancelPayment() (string, error) {
	return "", nil
}

func (p *Payment) ListPayments() (string, error) {
	return "", nil
}