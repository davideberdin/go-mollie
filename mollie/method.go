package mollie

type Method struct {
	Client Client
}

func (m *Method) GetMethod() (string, error) {
	return "", nil
}

func (m *Method) ListMethods() (string, error) {
	return "", nil
}
