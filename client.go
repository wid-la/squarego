package squarego

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// Service ...
type Service interface {
	GetCustomerByID(customerID string) (Customer, error)
	UpdateCustomerByID(customerID string, object Customer) (Customer, error)

	GetOrderByID(locationID, orderID string) (Order, error)
}

// NewService ...
func NewService(
	endpoint string,
	token string,
	version string,
) Service {

	return &service{
		Client:   &http.Client{},
		Endpoint: endpoint,
		Token:    token,
		Version:  version,
	}
}

type service struct {
	Client   *http.Client
	Endpoint string
	Token    string
	Version  string
}

func (svc *service) createRequest(method, ressource string, data []byte) (*http.Response, error) {
	var body io.Reader
	if data != nil {
		body = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method,
		fmt.Sprintf("%s/%s", svc.Endpoint, ressource),
		body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Square-Version", svc.Version)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", svc.Token))
	req.Header.Set("Accept", "application/json")

	resp, err := svc.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
