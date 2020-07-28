package squarego

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Payment ...
type Payment struct {
	ID            *string       `json:"id"`
	CreatedAt     *string       `json:"created_at"`
	UpdatedAt     *string       `json:"updated_at"`
	Status        *string       `json:"status"`
	SourceType    *string       `json:"source_type"`
	LocationID    *string       `json:"location_id"`
	OrderID       *string       `json:"order_id"`
	ReceiptNumber *string       `json:"receipt_number"`
	ReceiptURL    *string       `json:"receipt_url"`
	ReferenceID   *string       `json:"reference_id"`
	Note          *string       `json:"note"`
	CustomerID    *string       `json:"customer_id"`
	Version       *int          `json:"version"`
	AmountMoney   PaymentAmount `json:"amount_money"`
	TotalMoney    PaymentAmount `json:"total_money"`
	CardDetails   CardDetail    `json:"card_details"`
}

// PaymentResponse ...
type PaymentResponse struct {
	Payments []Payment `json:"payments"`
}

// PaymentAmount ...
type PaymentAmount struct {
	Amount   *int    `json:"amount"`
	Currency *string `json:"currency"`
}

// CardDetail ...
type CardDetail struct {
	AuthResultCode       *string `json:"auth_result_code"`
	AvsStatus            *string `json:"avs_status"`
	CvvStatus            *string `json:"cvv_status"`
	EntryMethod          *string `json:"entry_method"`
	StatementDescription *string `json:"statement_description"`
	Status               *string `json:"status"`
	Card                 Card    `json:"card"`
}

// Card ...
type Card struct {
	Bin         *string `json:"bin"`
	CardBrand   *string `json:"card_brand"`
	CardType    *string `json:"card_type"`
	ExpMonth    *int    `json:"exp_month"`
	ExpYear     *int    `json:"exp_year"`
	Fingerprint *string `json:"fingerprint"`
	LastFour    *string `json:"last_4"`
}

// GetRecentPayments ...
func (svc *service) GetRecentPayments(startDate time.Time) ([]Payment, error) {
	var paymentResp PaymentResponse

	paymentURL := fmt.Sprintf("payments?begin_time=%s&sort_order=ASC", startDate.Format(time.RFC3339))
	resp, err := svc.createRequest(http.MethodGet, paymentURL, nil)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &paymentResp)

	return paymentResp.Payments, nil
}
