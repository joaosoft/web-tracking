package web_tracking

import "encoding/json"

type TrackingRequest struct {
	Tracking struct {
		Category  *string          `json:"category" validate:"notzero"`
		Action    *string          `json:"action" validate:"notzero"`
		Label     *string          `json:"label"`
		Value     *int64           `json:"value"`
		Viewer    *string          `json:"viewer"`
		Viewed    *string          `json:"viewed"`
		Latitude  *float64         `json:"latitude"`
		Longitude *float64         `json:"longitude"`
		Street    *string          `json:"street"`
		MetaData  *json.RawMessage `json:"meta_data"`
	} `json:"tracking"`
}
