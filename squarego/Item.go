package squarego

// Item ...
type Item struct {
	UID            *string       `json:"uid"`
	Name           *string       `json:"name"`
	Quantity       *string       `json:"quantity"`
	BasePriceMoney PaymentAmount `json:"base_price_money"`
	TotalMoney     PaymentAmount `json:"total_money"`
}
