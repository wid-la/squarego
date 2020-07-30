package squarego

// Tender ...
type Tender struct {
	ID            *string `json:"id"`
	LocationID    *string `json:"location_id"`
	TransactionID *string `json:"transaction_id"`
	CustomerID    *string `json:"customer_id"`
}
