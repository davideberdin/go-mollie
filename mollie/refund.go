package mollie

type Refund struct {
	Client Client
}

func (r *Refund) CreateRefund() (string, error) {
	return "", nil
}

func (r *Refund) GetRefund() (string, error) {
	return "", nil
}

func (r *Refund) CancelRefund() (string, error) {
	return "", nil
}

func (r *Refund) ListRefunds() (string, error) {
	return "", nil
}