package squarego

import "time"

// Event ...
type Event struct {
	MerchantID *string   `json:"merchant_id"`
	Type       *string   `json:"type"`
	EventID    *string   `json:"event_id"`
	CreatedAt  *string   `json:"created_at"`
	Data       EventData `json:"data"`
}

// EventData ...
type EventData struct {
	ID     *string                `json:"id"`
	Type   *string                `json:"type"`
	Object map[string]interface{} `json:"object"`
}

// ConvertDateTime ...
func ConvertDateTime(date string) time.Time {
	layout := "2006-01-02T15:04:05.999Z"
	res, _ := time.Parse(layout, date)

	return res
}
