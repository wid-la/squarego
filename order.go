package squarego

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// OrderRequest ...
type OrderRequest struct {
	OrderIDs []string `json:"order_ids"`
}

// OrderResponse ...
type OrderResponse struct {
	Orders []Order `json:"orders"`
}

// Order ...
type Order struct {
	ID          *string       `json:"id"`
	LocationID  *string       `json:"location_id"`
	ReferenceID *string       `json:"reference_id"`
	CustomerID  *string       `json:"customer_id"`
	Items       []Item        `json:"line_items"`
	TotalMoney  PaymentAmount `json:"total_money"`
	Tenders     []Tender      `json:"tenders"`
}

// GetOrderByID ...
func (svc *service) GetOrderByID(locationID, orderID string) (Order, error) {
	var orderResp OrderResponse

	request := OrderRequest{
		OrderIDs: []string{orderID},
	}

	data, err := json.Marshal(request)
	if err != nil {
		return Order{}, err
	}

	resp, err := svc.createRequest(http.MethodPost, fmt.Sprintf("locations/%s/orders/batch-retrieve", locationID), data)
	if err != nil {
		return Order{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Order{}, err
	}

	err = json.Unmarshal(b, &orderResp)
	if len(orderResp.Orders) > 0 {
		return orderResp.Orders[0], nil
	}

	return Order{}, nil
}
