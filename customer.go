package squarego

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CustomerResponse ...
type CustomerResponse struct {
	Customer Customer `json:"customer"`
}

// Customer ...
type Customer struct {
	ID           *string `json:"id"`
	CreatedAt    *string `json:"created_at"`
	UpdatedAt    *string `json:"updated_at"`
	GivenName    *string `json:"given_name"`
	FamilyName   *string `json:"family_name"`
	EmailAddress *string `json:"email_address"`
	PhoneNumber  *string `json:"phone_number"`
	ReferenceID  *string `json:"reference_id"`
	Note         *string `json:"note"`
}

// GetCustomerByID ...
func (svc *service) GetCustomerByID(customerID string) (Customer, error) {
	var customerResp CustomerResponse

	resp, err := svc.createRequest(http.MethodGet, fmt.Sprintf("customers/%s", customerID), nil)
	if err != nil {
		return Customer{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Customer{}, err
	}

	err = json.Unmarshal(b, &customerResp)

	return customerResp.Customer, nil
}

// UpdateCustomerByID ...
func (svc *service) UpdateCustomerByID(customerID string, object Customer) (Customer, error) {
	var customerResp CustomerResponse

	data, err := json.Marshal(object)
	if err != nil {
		return object, err
	}

	resp, err := svc.createRequest(http.MethodPut, fmt.Sprintf("customers/%s", customerID), data)
	if err != nil {
		return Customer{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Customer{}, err
	}

	err = json.Unmarshal(b, &customerResp)

	return customerResp.Customer, nil
}
